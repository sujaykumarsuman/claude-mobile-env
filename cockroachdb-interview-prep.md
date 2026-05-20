# Cockroach Labs (CRL) — Backend MTS Interview Prep

Target rounds:
- **Thursday — DSA round** (1 hour, live coding, LeetCode Medium-level base + extension)
- **Friday — Choose Your Own System Design (CYOSD)** (1 hour, you pick & defend a system you built)

---

## Table of Contents

1. [How CRL evaluates these two rounds](#how-crl-evaluates-these-two-rounds)
2. [DSA Round — Difficulty Calibration](#dsa-round--difficulty-calibration)
3. [Language Choice](#language-choice)
4. [Tier 1 — Reported CRL LeetCode Questions](#tier-1--reported-crl-leetcode-questions)
5. [Tier 2 — Same-Pattern Problems](#tier-2--same-pattern-problems)
6. [Full Question List with Follow-Ups (44 problems)](#full-question-list-with-follow-ups)
7. [Concurrency Extension Vocab Card (C++)](#concurrency-extension-vocab-card-c)
8. [Three Rehearsal Prompts — Thread-Safe Extensions](#three-rehearsal-prompts--thread-safe-extensions)
9. [Choose Your Own System Design — Prep Plan](#choose-your-own-system-design-prep-plan)
10. [Two-Day Study Plan](#two-day-study-plan)
11. [Behavioral Stories to Prepare](#behavioral-stories-to-prepare)
12. [Sources](#sources)

---

## How CRL evaluates these two rounds

**DSA round (Thursday):** ~LeetCode Medium, often relevant to DB/distributed-systems primitives (graphs, intervals, parsing, tree/grid). Heavy emphasis on:

- Restating the problem before coding
- Sketching the approach on paper first
- Thinking aloud
- Clean, extensible code structure
- Handling an extension/follow-up they bolt on mid-round (~70% of rounds get this)

Scoring axes: correctness, debugging, algorithm choice, data structure choice, code structure, design sense.

**Choose Your Own System Design (Friday):** You are told the prompt in advance and bring a deep design *you* have worked on. The interviewer plays student. You must own it end-to-end: requirements → data model → APIs → storage → failure modes → scaling → operational concerns. Pick something you can defend at the level of "why this trade-off and not that one."

---

## DSA Round — Difficulty Calibration

From ~19 reported CRL coding problems across InterviewSolver / Algo.monster / Glassdoor:

| Difficulty | Share | Examples |
|---|---|---|
| Easy | ~20% | Same Tree, Merge Strings Alternately, Missing Ranges, Check If N and Its Double Exist |
| Medium | ~70% | String Compression, Basic Calculator II, Generate Parentheses, Kth Smallest in Sorted Matrix, Construct Quad Tree, Design Tic-Tac-Toe, Reveal Cards, Distinct Islands |
| Hard | ~10% | Making a Large Island, Number of Islands II |

Glassdoor difficulty rating: **3.1–3.4 / 5** — harder than average, not FAANG-Hard territory.

### The real difficulty is the extension

CRL's interview philosophy (from their blog): *"more representative of our day-to-day work — where we modify existing code to add features or handle changed requirements."*

So the realistic shape of the round is:

1. **Base problem (Medium, 25–30 min):** solve, test, complexity.
2. **Extension #1 (+Medium, 15–20 min):** "now support deletes" / "now it's a stream" / "now k changes" / "now there are concurrent writers."
3. **Extension #2 if time:** can push into Hard territory.

A starting problem that looks Medium (e.g. Number of Islands) can land you on Hard (Number of Islands II — Union-Find) by minute 35.

### What CRL weighs more vs less than FAANG

| Weighs more | Weighs less |
|---|---|
| Clean, extensible code structure | Obscure DP |
| Correct edge-case handling without prompting | Hard graph algorithms (Dijkstra/Floyd/MST rarely reported) |
| Talking through trade-offs while coding | Bit manipulation tricks |
| Picking the right data structure first try | Algorithmic cleverness for its own sake |

---

## Language Choice

CRL's own EngineeringExercises doc explicitly says: **use the language you know best, not Go just because the role is Go.** They are not scoring on language.

| Your fluency | Recommended |
|---|---|
| Strong C++, decent Go | **C++.** STL velocity (priority_queue, set, multiset, map, unordered_map) is a 15–20% speed win on heap/sweep problems. CRL doesn't care. |
| Strong Go, decent C++ | **Go.** Pre-memorize the heap boilerplate and you neutralize the main downside. |
| Strong Python | Fine. heapq, bisect, collections kit are great for these problems. |

**Do not language-shop mid-round.** Concurrency extensions on the DSA round are rare and the interviewer usually wants verbal reasoning, not a working goroutine program. See the [C++ concurrency vocab card](#concurrency-extension-vocab-card-c) below.

---

## Tier 1 — Reported CRL LeetCode Questions

These are the questions actually reported by candidates. Drill these first.

🔒 = LeetCode Premium-locked.

### Graphs / Grids (highest CRL signal)
- [200. Number of Islands](https://leetcode.com/problems/number-of-islands/) — Medium
- [305. Number of Islands II](https://leetcode.com/problems/number-of-islands-ii/) — Hard 🔒
- [694. Number of Distinct Islands](https://leetcode.com/problems/number-of-distinct-islands/) — Medium 🔒
- [827. Making A Large Island](https://leetcode.com/problems/making-a-large-island/) — Hard

### Strings / Parsing / Stack
- [443. String Compression](https://leetcode.com/problems/string-compression/) — Medium
- [227. Basic Calculator II](https://leetcode.com/problems/basic-calculator-ii/) — Medium
- [22. Generate Parentheses](https://leetcode.com/problems/generate-parentheses/) — Medium
- [1297. Maximum Number of Occurrences of a Substring](https://leetcode.com/problems/maximum-number-of-occurrences-of-a-substring/) — Medium
- [1768. Merge Strings Alternately](https://leetcode.com/problems/merge-strings-alternately/) — Easy

### Heap / Top-K / Matrix
- [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) — Medium
- [378. Kth Smallest Element in a Sorted Matrix](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/) — Medium

### Trees / Linked Lists
- [100. Same Tree](https://leetcode.com/problems/same-tree/) — Easy
- [427. Construct Quad Tree](https://leetcode.com/problems/construct-quad-tree/) — Medium
- [708. Insert into a Sorted Circular Linked List](https://leetcode.com/problems/insert-into-a-sorted-circular-linked-list/) — Medium 🔒
- [339. Nested List Weight Sum](https://leetcode.com/problems/nested-list-weight-sum/) — Medium 🔒

### Design + Misc
- [348. Design Tic-Tac-Toe](https://leetcode.com/problems/design-tic-tac-toe/) — Medium 🔒
- [950. Reveal Cards In Increasing Order](https://leetcode.com/problems/reveal-cards-in-increasing-order/) — Medium
- [163. Missing Ranges](https://leetcode.com/problems/missing-ranges/) — Easy 🔒
- [1346. Check If N and Its Double Exist](https://leetcode.com/problems/check-if-n-and-its-double-exist/) — Easy

### Mentioned by AlgoDaily (less canonical)
- [206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/) — Easy
- [283. Move Zeroes](https://leetcode.com/problems/move-zeroes/) — Easy
- [2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) — Medium
- [606. Construct String from Binary Tree](https://leetcode.com/problems/construct-string-from-binary-tree/) — Easy
- [404. Sum of Left Leaves](https://leetcode.com/problems/sum-of-left-leaves/) — Easy *(mirror of "Sum Right Side Leaves")*
- [258. Add Digits](https://leetcode.com/problems/add-digits/) — Easy
- [561. Array Partition](https://leetcode.com/problems/array-partition/) — Easy
- [509. Fibonacci Number](https://leetcode.com/problems/fibonacci-number/) — Easy

---

## Tier 2 — Same-Pattern Problems

Same buckets CRL pulls from. Strong candidates for extensions or close variants.

### Union-Find
- [721. Accounts Merge](https://leetcode.com/problems/accounts-merge/) — Medium
- [684. Redundant Connection](https://leetcode.com/problems/redundant-connection/) — Medium
- [547. Number of Provinces](https://leetcode.com/problems/number-of-provinces/) — Medium

### BFS on grid
- [286. Walls and Gates](https://leetcode.com/problems/walls-and-gates/) — Medium 🔒
- [994. Rotting Oranges](https://leetcode.com/problems/rotting-oranges/) — Medium

### Intervals
- [56. Merge Intervals](https://leetcode.com/problems/merge-intervals/) — Medium
- [57. Insert Interval](https://leetcode.com/problems/insert-interval/) — Medium
- [253. Meeting Rooms II](https://leetcode.com/problems/meeting-rooms-ii/) — Medium 🔒
- [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/) — Medium

### Heap / K-way merge
- [23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/) — Hard
- [295. Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/) — Hard

### Stack / Parsing
- [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) — Easy
- [394. Decode String](https://leetcode.com/problems/decode-string/) — Medium
- [224. Basic Calculator](https://leetcode.com/problems/basic-calculator/) — Hard

### Trees
- [98. Validate BST](https://leetcode.com/problems/validate-binary-search-tree/) — Medium
- [236. Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/) — Medium
- [297. Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/) — Hard

### Linked List
- [21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) — Easy

### Design
- [146. LRU Cache](https://leetcode.com/problems/lru-cache/) — Medium — **must know cold**
- [460. LFU Cache](https://leetcode.com/problems/lfu-cache/) — Hard
- [706. Design HashMap](https://leetcode.com/problems/design-hashmap/) — Easy
- [341. Flatten Nested List Iterator](https://leetcode.com/problems/flatten-nested-list-iterator/) — Medium 🔒

### Two-pointer / array hygiene
- [88. Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/) — Easy

### Official LeetCode CRL company tag
[leetcode.com/company/cockroach-labs/](https://leetcode.com/company/cockroach-labs/) — Premium-only, frequency-sorted, last-6-months filter is the highest-signal view.

---

## Full Question List with Follow-Ups

For each problem: **base idea** + **3 realistic CRL-style extensions**. CRL's extension patterns: *add a feature, change input model, make it concurrent, scale it, parameterize it.*

Drill format: solve the base, then for each follow-up, write the modified solution within 10 minutes.

---

### GROUP A — Graphs / Grids (HIGHEST CRL SIGNAL)

#### 1. Number of Islands — [LC 200, Medium](https://leetcode.com/problems/number-of-islands/)
**Core:** DFS/BFS over grid, mark visited.
- **F1:** Grid is too large to fit in memory; comes as a stream of rows. Running answer?
- **F2:** Support `addLand(r,c)` updates. *(Becomes Islands II — see #2.)*
- **F3:** Count islands by **shape** — two islands with same shape (rotation/reflection allowed) count as one type.

#### 2. Number of Islands II — [LC 305, Hard](https://leetcode.com/problems/number-of-islands-ii/) 🔒
**Core:** Union-Find with path compression + union by rank; decrement count on successful union.
- **F1:** Support `removeLand(r,c)` too. (DSU doesn't support deletion — rebuild periodically, or Link-Cut trees.)
- **F2:** Multi-threaded `addLand`.
- **F3:** Return not just the count but the **size of each island**. How do you maintain sizes under union?

#### 3. Number of Distinct Islands — [LC 694, Medium](https://leetcode.com/problems/number-of-distinct-islands/) 🔒
**Core:** DFS with normalized shape encoding (relative coordinates or path string).
- **F1:** Equivalent under **rotation and reflection** (LC 711 — Hard).
- **F2:** Stream of new lands; maintain count of distinct shapes seen so far.
- **F3:** Return the **largest** distinct island for each shape class.

#### 4. Making a Large Island — [LC 827, Hard](https://leetcode.com/problems/making-a-large-island/)
**Core:** Label connected components with IDs and sizes; for each 0-cell, sum sizes of unique neighbor IDs + 1.
- **F1:** Allowed to flip up to **K** zeros, not just 1. Approximation acceptable.
- **F2:** Cells have a *cost* to flip; budget B; maximize island size.
- **F3:** Online: a sequence of "what if I flip (r,c)?" queries — preprocess for O(1) per query.

---

### GROUP B — Strings / Parsing / Stack

#### 5. String Compression — [LC 443, Medium](https://leetcode.com/problems/string-compression/)
**Core:** Two pointers, in-place write.
- **F1:** Decompress back to original. Handle multi-digit counts.
- **F2:** Recursive compression: `"aaabbbaaa"` → `"a3b3a3"`, also detect `"abab"` → `"(ab)2"`.
- **F3:** Streaming input; emit compressed chars as soon as a run ends. Bound memory.

#### 6. Basic Calculator II — [LC 227, Medium](https://leetcode.com/problems/basic-calculator-ii/)
**Core:** Stack; defer `+/-`, apply `*//` immediately against top.
- **F1:** Add **parentheses** support (becomes LC 224).
- **F2:** Add **unary minus** and **floating point**.
- **F3:** Parse and return an **AST**, not just the result. Then evaluate. (Sets up "now optimize repeated evaluation" — query-planner question.)

#### 7. Generate Parentheses — [LC 22, Medium](https://leetcode.com/problems/generate-parentheses/)
**Core:** Backtracking with open/close counters.
- **F1:** Generate only the **kth** valid string in lex order without enumerating all (Catalan ranking).
- **F2:** Now 3 bracket types `(){}[]`. Generate all valid balanced strings of length 2n.
- **F3:** Given a string with `?` chars, return count of valid bracket fillings.

#### 8. Max Occurrences of a Substring — [LC 1297, Medium](https://leetcode.com/problems/maximum-number-of-occurrences-of-a-substring/)
**Core:** Sliding window of size `minSize` only; count via hashmap.
- **F1:** Constraint changes to "exactly K distinct chars."
- **F2:** Streaming text; maintain top-K most frequent substrings of size m. (Count-Min Sketch territory.)
- **F3:** Find substring of length `[minSize, maxSize]` with **highest occurrences per character**.

#### 9. Merge Strings Alternately — [LC 1768, Easy](https://leetcode.com/problems/merge-strings-alternately/)
**Core:** Two pointers.
- **F1:** Merge N strings round-robin, not just 2.
- **F2:** Skip vowels while alternating.
- **F3:** Make it a generator/iterator yielding one char at a time (streaming).

#### 10. Decode String — [LC 394, Medium](https://leetcode.com/problems/decode-string/)
**Core:** Two stacks (count + accumulated string) or recursive parse.
- **F1:** Support nested brackets with negative repeats meaning "reverse."
- **F2:** Don't materialize the full string — answer queries like `charAt(i)` in O(log n) via the parse tree.
- **F3:** Validate input first; reject malformed strings with exact error position.

#### 11. Valid Parentheses — [LC 20, Easy](https://leetcode.com/problems/valid-parentheses/)
**Core:** Stack.
- **F1:** Return the **longest valid prefix**. Then longest valid substring (LC 32).
- **F2:** Allow wildcard `*` matching `(`, `)`, or empty (LC 678).
- **F3:** Streaming version: O(1) per char with rolling validity check.

---

### GROUP C — Heap / Top-K / Matrix

#### 12. Kth Largest Element in Array — [LC 215, Medium](https://leetcode.com/problems/kth-largest-element-in-an-array/)
**Core:** Min-heap of size k OR Quickselect.
- **F1:** Stream version — maintain top-k as elements arrive (LC 703).
- **F2:** Kth largest **across multiple sorted arrays**.
- **F3:** k can change between queries; support `add(x)` and `kthLargest(k)` efficiently.

#### 13. Kth Smallest in Sorted Matrix — [LC 378, Medium](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/)
**Core:** Min-heap of row-frontiers (O(k log n)) OR binary search on value (O(n log(max-min))).
- **F1:** Matrix is enormous; binary-search-on-value variant only.
- **F2:** Find the kth smallest **sum** of one element per row.
- **F3:** Support updates: change one cell, then ask kth smallest again.

#### 14. Merge K Sorted Lists — [LC 23, Hard](https://leetcode.com/problems/merge-k-sorted-lists/)
**Core:** Min-heap of head pointers.
- **F1:** Lists are huge; you only get iterators. Same algorithm, defend memory.
- **F2:** Multi-threaded: each list on a separate thread. Merge into a single output stream.
- **F3:** Find smallest range covering at least one number from each list (LC 632).

---

### GROUP D — Trees

#### 15. Same Tree — [LC 100, Easy](https://leetcode.com/problems/same-tree/)
**Core:** Recursive structural compare.
- **F1:** Subtree of Another Tree (LC 572).
- **F2:** Trees equivalent if children can be in any order (unordered tree equality).
- **F3:** Trees too deep for recursion — iterative version with explicit stack.

#### 16. Construct Quad Tree — [LC 427, Medium](https://leetcode.com/problems/construct-quad-tree/)
**Core:** Recursive divide; merge when all four children agree.
- **F1:** Implement **intersect** of two quadtrees (LC 558).
- **F2:** Support point updates: flip cell (r,c); rebalance structure.
- **F3:** Range query: count of 1s in subrectangle in O(log n).

#### 17. Validate BST — [LC 98, Medium](https://leetcode.com/problems/validate-binary-search-tree/)
**Core:** In-order strictly increasing, or pass (min, max) bounds down.
- **F1:** Allow duplicates on the left subtree only.
- **F2:** Find the two **swapped nodes** that broke the BST property (LC 99).
- **F3:** Validate incrementally as nodes are inserted (rolling invariant).

#### 18. Lowest Common Ancestor — [LC 236, Medium](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)
**Core:** Recursive: if both found in different subtrees, current is LCA.
- **F1:** Each node has a `parent` pointer — solve in O(h) time, O(1) extra space.
- **F2:** Many queries on same tree — preprocess with binary lifting for O(log n) per query.
- **F3:** LCA of K nodes, not just 2.

#### 19. Serialize and Deserialize Binary Tree — [LC 297, Hard](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/)
**Core:** Pre-order with null markers.
- **F1:** Optimize for size: compress repeated subtrees (DAG encoding).
- **F2:** Stream-friendly serialization — produce bytes as you traverse; consume bytes as they arrive.
- **F3:** Backward compatibility: old serialized blobs must still deserialize after adding a node-type field.

#### 20. Sum of Left Leaves — [LC 404, Easy](https://leetcode.com/problems/sum-of-left-leaves/)
**Core:** DFS with a flag for "am I a left child."
- **F1:** Sum **only leaves at the deepest level** (LC 1302).
- **F2:** Sum of all leaves grouped by depth (return a map).
- **F3:** Maximum path sum that starts and ends at leaves only.

#### 21. Construct String from Binary Tree — [LC 606, Easy](https://leetcode.com/problems/construct-string-from-binary-tree/)
**Core:** Pre-order with parentheses, rules for omitting empty parens.
- **F1:** Inverse — parse the string back to a tree (LC 536).
- **F2:** Pretty-print with indentation instead of parens.
- **F3:** Stream output; bound output buffer to N bytes; truncate gracefully.

---

### GROUP E — Linked Lists

#### 22. Reverse Linked List — [LC 206, Easy](https://leetcode.com/problems/reverse-linked-list/)
**Core:** Three-pointer iterative or recursive.
- **F1:** Reverse in **groups of k** (LC 25, Hard).
- **F2:** Reverse only the **even-positioned nodes**.
- **F3:** Reverse between positions m and n in one pass (LC 92).

#### 23. Add Two Numbers — [LC 2, Medium](https://leetcode.com/problems/add-two-numbers/)
**Core:** Iterate with carry.
- **F1:** Digits stored **most-significant-first** (LC 445). Stack or reverse.
- **F2:** Multiply, not add.
- **F3:** Add in arbitrary base (e.g., base 7, base 16).

#### 24. Insert into Sorted Circular Linked List — [LC 708, Medium](https://leetcode.com/problems/insert-into-a-sorted-circular-linked-list/) 🔒
**Core:** Walk until the wrap-around point or the slot between two nodes.
- **F1:** Delete a value; handle head-deletion.
- **F2:** Thread-safe inserts.
- **F3:** Support `kthLargest(k)` after each insert.

#### 25. Move Zeroes — [LC 283, Easy](https://leetcode.com/problems/move-zeroes/)
**Core:** Two pointers, in-place.
- **F1:** Move all instances of a given value to the end, preserving relative order.
- **F2:** Stable partition — even-indexed first, odd-indexed after.
- **F3:** Same problem on a **doubly linked list**.

#### 26. Merge Two Sorted Lists — [LC 21, Easy](https://leetcode.com/problems/merge-two-sorted-lists/)
**Core:** Dummy head + two pointers.
- **F1:** Merge **k** lists (LC 23).
- **F2:** Deduplicate during merge.
- **F3:** Merge two sorted iterators with no node mutation allowed; return a new list.

---

### GROUP F — Intervals / Sweep

#### 27. Merge Intervals — [LC 56, Medium](https://leetcode.com/problems/merge-intervals/)
**Core:** Sort by start; sweep.
- **F1:** Streaming — intervals arrive online; maintain merged set with fast insert.
- **F2:** Intervals are weighted; report max total weight at any point.
- **F3:** Multi-dimensional rectangles, not 1D intervals.

#### 28. Insert Interval — [LC 57, Medium](https://leetcode.com/problems/insert-interval/)
**Core:** Three-phase sweep: before / overlap-merge / after.
- **F1:** Many insertions — use a balanced BST keyed by start.
- **F2:** Insert a "subtract" interval that punches a hole.
- **F3:** Track count of overlapping intervals at each point (analogous to MVCC version count).

#### 29. Meeting Rooms II — [LC 253, Medium](https://leetcode.com/problems/meeting-rooms-ii/) 🔒
**Core:** Min-heap of end times, OR sweep line with +1/-1 events.
- **F1:** Return the **schedule** (which meeting → which room), not just the count.
- **F2:** Rooms have **capacities**; minimize rooms.
- **F3:** Streaming meetings; answer "rooms needed so far" after each arrival.

#### 30. Missing Ranges — [LC 163, Easy](https://leetcode.com/problems/missing-ranges/) 🔒
**Core:** Linear scan with previous pointer.
- **F1:** Stream of sorted-but-sparse numbers; emit missing ranges as you go without buffering.
- **F2:** Numbers can repeat; same logic but skip dupes.
- **F3:** Given two sorted lists, find ranges present in A but missing in B.

---

### GROUP G — Union-Find / Connectivity

#### 31. Accounts Merge — [LC 721, Medium](https://leetcode.com/problems/accounts-merge/)
**Core:** Union-Find with email→account mapping.
- **F1:** Dedupe by **phone number** in addition to email.
- **F2:** Streaming — new accounts arrive; maintain merged view online.
- **F3:** Return the merge graph — *why* were two accounts merged (which shared keys)?

#### 32. Redundant Connection — [LC 684, Medium](https://leetcode.com/problems/redundant-connection/)
**Core:** Union-Find; first edge that connects already-connected nodes is redundant.
- **F1:** Directed version (LC 685, Hard).
- **F2:** Return **all** redundant edges.
- **F3:** Edges have weights; remove the one whose removal least increases shortest paths.

#### 33. Number of Provinces — [LC 547, Medium](https://leetcode.com/problems/number-of-provinces/)
**Core:** Count of DSU components.
- **F1:** Dynamic edges: support `connect(u,v)` and `query()` for count.
- **F2:** Now also support `disconnect(u,v)`. (Harder — offline trick or Link-Cut.)
- **F3:** Weighted: each province has a "rank" = sum of node weights; return ranks.

---

### GROUP H — Design

#### 34. Design Tic-Tac-Toe — [LC 348, Medium](https://leetcode.com/problems/design-tic-tac-toe/) 🔒
**Core:** O(1) per move: maintain row/col/diag counters per player.
- **F1:** Generalize to **N-in-a-row on M×M board** (Gomoku).
- **F2:** Three-player version.
- **F3:** Undo last move in O(1).

#### 35. LRU Cache — [LC 146, Medium](https://leetcode.com/problems/lru-cache/) **MUST KNOW COLD**
**Core:** Doubly-linked list + hashmap.
- **F1:** **LFU Cache** (LC 460, Hard).
- **F2:** **TTL** — entries expire after N seconds; lazy or active eviction?
- **F3:** Thread-safe and high-throughput. *(See rehearsal section.)*

#### 36. Design HashMap — [LC 706, Easy](https://leetcode.com/problems/design-hashmap/)
**Core:** Bucket array + chaining.
- **F1:** Implement **resize** when load factor > threshold.
- **F2:** Thread-safe via per-bucket locks. *(See rehearsal section.)*
- **F3:** Implement **iteration** with stable order under concurrent writes.

#### 37. Flatten Nested List Iterator — [LC 341, Medium](https://leetcode.com/problems/flatten-nested-list-iterator/) 🔒
**Core:** Stack of iterators; `hasNext` peels nesting lazily.
- **F1:** Support `prev()` as well as `next()`.
- **F2:** Nesting can be infinite (lazy generators); never materialize all.
- **F3:** Multi-threaded consumers; concurrent `next()` calls.

#### 38. Nested List Weight Sum — [LC 339, Medium](https://leetcode.com/problems/nested-list-weight-sum/) 🔒
**Core:** DFS with depth, or BFS by level.
- **F1:** **Inverse depth** weighting (LC 364).
- **F2:** Add `sumByDepth()` query returning per-depth breakdown.
- **F3:** Stream the structure token-by-token; compute weighted sum online.

---

### GROUP I — Misc Reported Questions

#### 39. Reveal Cards in Increasing Order — [LC 950, Medium](https://leetcode.com/problems/reveal-cards-in-increasing-order/)
**Core:** Simulate with deque of indices; pop+rotate.
- **F1:** **Reverse** the process — given the reveal order, reconstruct the original deck.
- **F2:** Variable rotation step k.
- **F3:** Multiple decks interleaved; find the global reveal order.

#### 40. Check If N and Its Double Exist — [LC 1346, Easy](https://leetcode.com/problems/check-if-n-and-its-double-exist/)
**Core:** Hashset; for each x check x*2 or x/2.
- **F1:** Generalize to "x and k*x both exist" for arbitrary k (handle negatives, zeros).
- **F2:** Return **all such pairs**, not just existence.
- **F3:** Streaming — answer "does the double exist so far" after each arrival.

#### 41. Add Digits / Sum Digits Until One — [LC 258, Easy](https://leetcode.com/problems/add-digits/)
**Core:** O(1) closed form via digital root: `1 + (n-1) % 9`. Also know the loop solution.
- **F1:** Same problem in base B.
- **F2:** Return the **path** of intermediate sums.
- **F3:** Given a count target k, find the smallest n that takes exactly k iterations.

#### 42. Array Partition / "Max of Min Pairs" — [LC 561, Easy](https://leetcode.com/problems/array-partition/)
**Core:** Sort, sum every even-indexed element.
- **F1:** Prove correctness of pair-adjacent-after-sort. (Exchange argument.)
- **F2:** Generalize to triples: maximize sum of min of each triple.
- **F3:** Stream of numbers; maintain running answer.

#### 43. Fibonacci — [LC 509, Easy](https://leetcode.com/problems/fibonacci-number/)
**Core:** O(n) iterative, or O(log n) matrix exponentiation.
- **F1:** Compute **nth Fibonacci mod M** for huge n (matrix exp).
- **F2:** **Pisano period** — count Fibonacci numbers divisible by k in a range.
- **F3:** Given a number, return its **Zeckendorf representation**.

#### 44. Decimal to Binary
**Core:** Repeated division by 2; collect remainders.
- **F1:** Convert between **arbitrary bases** (LC 504 base 7, LC 168 Excel column).
- **F2:** Handle **negative numbers** in two's complement.
- **F3:** Convert a **fractional decimal** to binary.

---

## Concurrency Extension Vocab Card (C++)

Spend 30 minutes internalizing this. Covers ~95% of any DSA-round concurrency follow-up.

```cpp
#include <mutex>
#include <shared_mutex>
#include <atomic>
#include <condition_variable>
#include <thread>

std::mutex m;                    // exclusive
std::shared_mutex rw;            // reader-writer
std::atomic<int> counter{0};     // lock-free counter
std::condition_variable cv;

// Exclusive lock (RAII — auto-unlocks)
{
    std::lock_guard<std::mutex> lk(m);
    // critical section
}

// Reader-writer
{
    std::shared_lock<std::shared_mutex> lk(rw);   // many readers
    // read
}
{
    std::unique_lock<std::shared_mutex> lk(rw);   // one writer
    // write
}

// Atomic increment (no lock needed)
counter.fetch_add(1, std::memory_order_relaxed);

// Condition variable wait
std::unique_lock<std::mutex> lk(m);
cv.wait(lk, []{ return ready; });
cv.notify_one();   // or notify_all()
```

### Concepts to name out loud

- **Race condition** — unsynchronized concurrent access to shared state
- **Deadlock** — two locks acquired in inconsistent order
- **Lock granularity** — coarse (one big lock) vs. fine (per-bucket, per-row)
- **Reader-writer trade-off** — wins when reads ≫ writes
- **Lock-free / atomic** — for counters and flags; avoids contention but harder to reason about
- **Memory ordering** — `relaxed` for counters, `acquire/release` for handoff
- **Sharding** — split state into N independent locked partitions to reduce contention
- **Copy-on-write / immutability** — sometimes cleaner than locking

### Go equivalents (if asked to translate)
- `std::mutex` → `sync.Mutex`
- `std::shared_mutex` → `sync.RWMutex`
- `std::atomic<int>` → `sync/atomic.Int64`
- `std::thread` → `go func() { ... }()`
- `condition_variable` → channels + `sync.Cond`

---

## Three Rehearsal Prompts — Thread-Safe Extensions

### Prompt 1 — LRU Cache (LC 146)

**Extension:** *"Now multiple threads will call get and put concurrently. Make it thread-safe."*

**Model verbal answer (~90 sec):**

> "The non-obvious thing about LRU is that **`get` is not a read** — it mutates the list by splicing the accessed node to the front. So I can't use a reader-writer lock naively; both `get` and `put` are writers from the sync standpoint.
>
> **Simplest correct answer:** wrap both methods in a single `std::mutex` with `std::lock_guard`. Coarse but correct, and for an in-memory cache the critical section is microseconds.
>
> **If they push for better throughput:** shard the cache into N independent LRUs by `hash(key) % N`, each with its own mutex. Linear contention reduction. Trade-off: eviction is per-shard, not global.
>
> **If they push into lock-free:** wouldn't go there in an interview — research-level problem. Coarse lock + sharding is the right engineering call."

**Likely probes:**
- *"Why not shared_mutex?"* → `get` mutates the list; it's not really a read.
- *"What if a thread dies inside the critical section?"* → `lock_guard` is RAII; lock releases on stack unwind.
- *"Iterator invalidation?"* → With the mutex held, no — `std::list` iterators are stable across operations on other nodes.

**Wrong-answer trap:** `std::shared_mutex` with `get` under shared_lock. Two readers both splicing the list corrupts pointers.

---

### Prompt 2 — Design HashMap (LC 706)

**Extension:** *"Make it safe and high-throughput for many concurrent threads."*

**Model verbal answer (~90 sec):**

> "This is the textbook case for **fine-grained locking via bucket-level mutexes** — exactly what Java's `ConcurrentHashMap` does.
>
> ```cpp
> static constexpr int N = 1024;
> std::vector<std::list<std::pair<int,int>>> buckets{N};
> std::vector<std::shared_mutex> locks{N};
> ```
>
> Reads use `shared_lock` so many concurrent `get`s on the same bucket parallelize. Writes serialize, and only within the same bucket.
>
> **Hard part is resize.** Two options:
> 1. Stop-the-world: take all N locks in fixed order, rehash, release. Latency spike.
> 2. Incremental: keep two tables, lazily migrate buckets on access. No spike.
>
> For interview, option 1; mention option 2 as the production answer."

**Likely probes:**
- *"Why fixed N and not grow the lock array?"* → Lock arrays growing concurrently is itself a sync problem. Fix N at construction; locks are cheap.
- *"Bad hash function?"* → Distribution skews, one lock becomes hot, throughput collapses to single-threaded.
- *"Lock-free?"* → Cliff Click / folly exist but 1000+ lines, not interview-scope.

**Wrong-answer trap:** A single global mutex — defeats the whole point.

---

### Prompt 3 — Number of Islands II (LC 305)

**Extension:** *"Now `addLand` calls arrive from multiple threads. How do you make this safe and still correct?"*

**Model verbal answer (~90 sec):**

> "Union-Find is **tricky to parallelize** because `find` with path compression *mutates* parent pointers — so even a `find` is a write.
>
> **Simplest correct:** one `std::mutex` guarding the whole DSU. Every `addLand` takes the lock, does the up-to-4 unions, decrements `count`, releases.
>
> **Higher throughput** is genuinely hard because union of cells in different regions still touches shared roots after a few merges. Options:
> 1. Partition the grid into tiles with per-tile DSUs. Within-tile lands take the tile lock; cross-tile take a global lock.
> 2. Lock-free DSU (Anderson, Woelfel) — research literature, wouldn't implement in an hour.
>
> **Subtle:** the returned `count` must reflect state after this addLand and before any concurrent one. Read `count` inside the same critical section that mutated it."

**Likely probes:**
- *"shared_mutex?"* → `find` mutates via path compression; not a read. Disabling path compression makes `find` O(log n) read-only — workload-dependent.
- *"Per-cell locks?"* → Doesn't help; contention is on *roots*, not cells.
- *"Worst-case latency?"* → O(α(n)) per addLand, effectively constant. Coarse lock is fine until ~10M addLands/sec.

**Wrong-answer trap:** Locking only `union` but not `find`. Path-compression mutation corrupts the parent array under contention.

---

### Drill ritual for all three
1. Read the model answer twice, then close it.
2. Set a 90-second timer. Say the answer out loud as if to an interviewer. Record if possible.
3. Replay. Every "uh", every stall, every probe you couldn't answer — that's a gap.
4. Do all 3 in a row to build endurance.

Pattern across all three (this is what CRL scores):
- Coarse lock first → state correctness
- Identify the contention bottleneck
- Propose fine-grained or sharded approach
- Name the trade-off out loud
- Decline lock-free with a clear reason

---

## Choose Your Own System Design — Prep Plan

### Picking the topic

Pick a system that satisfies all four:
1. **You personally designed or led a major piece of it** (not "I worked on the team that…").
2. **Non-trivial distributed/consistency/scale dimension** — replication, sharding, queueing, idempotency, leader election, caching, schema migration.
3. **You can defend at least 3 design trade-offs** with the alternative you rejected and why.
4. **You can draw it on a whiteboard from memory in 5 minutes.**

If you have two candidates, pick the one with more **failure stories** — outages, regressions, rollbacks. CRL interviewers love probing what broke.

### Structure to deliver (45–50 min)

1. **Context (3 min):** business problem, who used it, scale numbers (QPS, data size, latency SLO), your specific role.
2. **Requirements (3 min):** functional + non-functional. State what you chose **not** to support.
3. **High-level architecture (5 min):** boxes & arrows. Name every box.
4. **Data model (5 min):** tables/schemas/keys. If sharded — sharding key and why. If consistency matters — isolation level and why.
5. **API / request path (5 min):** walk a read and a write end-to-end.
6. **Hard parts (10 min):** the 2–3 actually-difficult problems. **This is where you win or lose.** Examples:
   - Hot partition mitigation
   - Exactly-once vs at-least-once + idempotency keys
   - Failover / split brain
   - Schema migration without downtime
   - Backpressure under load spike
7. **Scale / failure modes (5 min):** what falls over first at 10× load, what we monitored, what we paged on.
8. **What you'd change with hindsight (3 min):** humility + senior judgment signal.

### Probes to pre-rehearse answers for

- "Why this database / queue / cache and not X?"
- "What's the consistency model? What anomalies are possible?"
- "How do you handle a node dying mid-write?"
- "How do you roll out a backwards-incompatible change?"
- "What happens at 10× traffic? 100×?"
- "Where's the single point of failure right now?"
- "How do you test this? How do you know it's correct in prod?"
- "What did you measure? What were the actual numbers?"
- **CRL-specific:** "If you replaced your storage layer with CockroachDB, what would simplify and what would get harder?" — one-paragraph answer here scores big.

### Red flags to avoid

- Choosing a system you only partially worked on. They **will** drill into the part you didn't own.
- Hand-waving numbers ("a lot of traffic"). Bring real numbers or honest estimates with the math.
- Defending every decision as optimal. Pre-pick 2 things you'd do differently.
- Not drawing. Whiteboard/screen-share a diagram even if remote — silence + words is worse than mediocre boxes.

### CockroachDB vocabulary to skim (30 min, high leverage)

Skim CRL architecture so you can speak their language if relevant:
- **Range** — contiguous key-space partition (default 512 MiB)
- **Raft** — consensus per range; leaseholder serves reads
- **MVCC** — versioned keys with HLC timestamps
- **Gateway node** — the node receiving the SQL connection
- **Distributed SQL (DistSQL)** — query execution pushed to data
- **Serializable isolation** — default; uses SSI

---

## Two-Day Study Plan

### Day 1 (Wednesday — before Thursday DSA round)

**Morning (3 hrs):** Groups A + B
- #1, #2, #4 (graphs/grids)
- #5, #6, #7, #8 (strings/parsing)
- One follow-up each.

**Afternoon (3 hrs):** Groups C + D
- #12, #13, #14 (heaps)
- #16, #18, #19 (trees)
- One follow-up each.

**Evening (1.5 hrs):** Group H design
- #34, #35, #36 — these are extension *magnets*. Do all three follow-ups for each.

### Day 2 (Thursday evening — between rounds, before Friday CYOSD)

**Thursday post-DSA decompress:** 1–2 hrs off.

**Then 3 hrs on CYOSD:**
- Write a one-page brief of your chosen system covering the 8 sections.
- Explain it out loud to a wall/voice memo. Listen back — every "um, basically" is a place you don't know the answer.
- Anticipate 8 probes. Write 2-sentence answers to each.
- Sketch the architecture diagram 3 times from blank until muscle memory.

**Friday morning:**
- Re-draw the diagram once.
- Re-read your trade-off list.
- Skim CockroachDB architecture vocabulary.

### Drill rules

1. **25 min budget per base problem, 10 min per follow-up.** If you blow the budget, read the editorial, code from memory, mark for re-attempt tomorrow.
2. **Narrate out loud** for every solve. The verbal pattern is muscle memory by Thursday or it isn't.
3. **Keep a "missed list."** Re-do only that list Friday morning.

---

## Behavioral Stories to Prepare

CRL screens for: collaboration, opinion-holding-loosely, debugging instinct, ownership of failure.

Have one crisp story (STAR format: Situation, Task, Action, Result) each for:

1. **A production incident you led the response to** — what broke, how you diagnosed, what you fixed, what you changed permanently.
2. **A design decision you reversed** — what you originally chose, what changed your mind, how you communicated the reversal.
3. **A disagreement with a senior engineer where you were right** — show conviction + diplomacy.
4. **A disagreement where you were wrong** — show ability to update.
5. **Something you shipped that mattered to a customer** — customer empathy + outcome focus.

Each story: 60–90 seconds spoken. Quantify outcomes ("reduced p99 from 400ms to 60ms"; "unblocked 12 customer accounts").

---

## Sources

- [Cockroach Labs Open Sourced Interview Process (GitHub)](https://github.com/cockroachlabs/open-sourced-interview-process)
- [EngineeringExercises.md](https://github.com/cockroachlabs/open-sourced-interview-process/blob/master/EngineeringExercises.md)
- [CRL blog: Updating the engineering interview](https://www.cockroachlabs.com/blog/updating-eng-interview/)
- [CRL Open Interview Process page](https://www.cockroachlabs.com/careers/open-interview/)
- [Glassdoor — Cockroach Labs SWE Interview](https://www.glassdoor.com/Interview/Cockroach-Labs-Software-Engineer-Interview-Questions-EI_IE1168502.0,14_KO15,32.htm)
- [Blind — Cockroach Lab Interview thread](https://www.teamblind.com/company/Cockroach-Lab/posts/cockroach-lab-interview)
- [Algo.monster CRL guide](https://algo.monster/interview-guides/cockroach-labs)
- [InterviewSolver — CRL question list](https://interviewsolver.com/interview-questions/cockroach-labs)
- [AlgoDaily — CRL question list](https://algodaily.com/companies/cockroach-labs)
- [LeetCode CRL company tag (Premium)](https://leetcode.com/company/cockroach-labs/)
