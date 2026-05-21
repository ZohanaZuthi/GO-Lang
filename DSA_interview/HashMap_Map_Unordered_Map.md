# HashMap / Map / Unordered Map — DSA Interview Notes

In C++, when people say **“hashmap”**, they usually mean:

- `unordered_map` → hash table based
- `map` → balanced binary search tree based

Both store:

```cpp
key -> value
```

Example:

```cpp
name -> marks
"a"  -> 5
10   -> 3
```

---

# 1. What Problem Does HashMap Solve?

Main purpose:

- Fast searching
- Fast insertion
- Fast deletion
- Frequency counting
- Mapping relations

Example:

```cpp
arr = [1,1,2,3,3,3]

1 -> 2 times
2 -> 1 time
3 -> 3 times
```

Instead of nested loops `O(n²)`,
hashmap does it in almost `O(n)`.

---

# 2. Two STL Structures

## A. `map`

```cpp
map<int,int> mp;
```

Implemented using:

- Red Black Tree (self-balancing BST)

Keys stay sorted automatically.

Example:

```cpp
3 -> x
1 -> y
2 -> z
```

Stored as:

```cpp
1 2 3
```

---

## B. `unordered_map`

```cpp
unordered_map<int,int> mp;
```

Implemented using:

- Hash Table

No ordering.

Example output order:

```cpp
2 1 3
```

or

```cpp
3 2 1
```

Random-like order.

---

# 3. Complexity Difference

| Operation         | `map`    | `unordered_map` |
| ----------------- | -------- | --------------- |
| Insert            | O(log n) | O(1) average    |
| Search            | O(log n) | O(1) average    |
| Delete            | O(log n) | O(1) average    |
| Ordered traversal | Yes      | No              |

---

# 4. WHY `map` is O(log n)

Because it uses a balanced BST.

A balanced BST height becomes:

h \approx \log_2 n

Searching in a tree depends on height.

Example:

For 16 elements:

```text
height ≈ log₂(16) = 4
```

So search takes around 4 steps.

That is why:

```text
map => O(log n)
```

---

# 5. WHY `unordered_map` is O(1)

It uses hashing.

The key is converted into an array index.

Example:

```cpp
key = 25
hash(25) = 5
```

Then directly go to bucket/index 5.

No tree traversal needed.

Direct access:

```text
O(1)
```

---

# 6. Then Why NOT Always Use unordered_map?

Because collisions happen.

---

# 7. Collision

Two keys may produce same hash.

Example:

```text
hash(15) = 5
hash(25) = 5
hash(35) = 5
```

All go to same bucket.

This is collision.

---

# 8. Worst Case Complexity

If all elements collide:

```text
unordered_map => O(n)
```

Because bucket becomes like a linked list.

So:

| Structure     | Average  | Worst    |
| ------------- | -------- | -------- |
| map           | O(log n) | O(log n) |
| unordered_map | O(1)     | O(n)     |

---

# 9. Why unordered_map Is Still Faster

Because collisions are usually rare.

Good hash functions distribute values well.

So average performance is excellent.

That is why competitive programmers usually use:

```cpp
unordered_map
```

for frequency counting.

---

# 10. Internal Idea of Hash Table

Think:

```text
index -> bucket
```

Like:

```text
0 -> []
1 -> [11]
2 -> [22, 32]
3 -> []
4 -> [14]
```

Bucket stores collided values.

---

# 11. Rehashing

When table becomes crowded:

- size increases
- elements redistributed

This is called:

```text
rehashing
```

Load factor controls this.

---

# 12. Load Factor

Formula:

\text{Load Factor} = \frac{\text{Number of Elements}}{\text{Number of Buckets}}

High load factor:

- more collisions
- slower performance

---

# 13. Frequency Counting

Most common hashmap use.

Example:

```cpp
vector<int> arr = {1,2,2,3,3,3};

unordered_map<int,int> freq;

for(int x : arr)
{
    freq[x]++;
}
```

Result:

```text
1 -> 1
2 -> 2
3 -> 3
```

Complexity:

```text
O(n)
```

---

# 14. Character Frequency

Very common in interviews.

```cpp
string s = "aabbccc";

unordered_map<char,int> freq;

for(char c : s)
{
    freq[c]++;
}
```

Result:

```text
a -> 2
b -> 2
c -> 3
```

---

# 15. Can Characters Be Keys?

YES.

```cpp
unordered_map<char,int>
```

Because internally chars are integers (ASCII).

Example:

```text
'a' = 97
'b' = 98
```

---

# 16. String as Key

YES.

```cpp
unordered_map<string,int>
```

Example:

```cpp
"apple" -> 5
"banana" -> 2
```

Hash function converts string into integer hash.

---

# 17. Pair as Key

You partially misunderstood this part.

## `map`

Works directly:

```cpp
map<pair<int,int>, int> mp;
```

Because pairs can be compared in BST ordering.

---

## `unordered_map`

Does NOT directly support pair.

This is because hashing for pair is not built-in.

This fails:

```cpp
unordered_map<pair<int,int>, int> mp;
```

Need custom hash function.

---

# 18. Custom Hash for Pair

Example:

```cpp
struct hash_pair
{
    size_t operator()(const pair<int,int>& p) const
    {
        return p.first ^ p.second;
    }
};

unordered_map<pair<int,int>, int, hash_pair> mp;
```

Interviewers love this topic.

---

# 19. Division Hashing

You mentioned:

```text
key % m
```

Correct.

Very common hash function.

Example:

```text
25 % 10 = 5
35 % 10 = 5
```

Collision happens.

---

# 20. How Character Hashing Works

Characters are integers internally.

Example:

```text
'a' % 10
97 % 10 = 7
```

So same hashing concept applies.

---

# 21. Common Interview Questions

## Q1: Difference Between map and unordered_map?

Answer:

| map        | unordered_map |
| ---------- | ------------- |
| BST        | Hash Table    |
| Ordered    | Unordered     |
| O(log n)   | O(1) average  |
| No hashing | Uses hashing  |

---

## Q2: When Use map?

Use when:

- sorted order needed
- lower_bound / upper_bound needed
- ordered traversal needed

Example:

```cpp
map<int,int>
```

---

## Q3: When Use unordered_map?

Use when:

- only fast lookup needed
- ordering not important

Usually faster.

---

# 22. lower_bound Advantage of map

Since `map` is sorted:

```cpp
mp.lower_bound(x)
```

is possible.

Not possible efficiently in `unordered_map`.

Huge interview point.

---

# 23. Ordered Traversal

```cpp
map<int,int> mp;

mp[3]=10;
mp[1]=20;
mp[2]=30;
```

Traversal:

```text
1 2 3
```

Automatic sorting.

---

# 24. Important STL Syntax

## Insert

```cpp
mp[key] = value;
```

---

## Frequency

```cpp
freq[x]++;
```

---

## Check Existence

```cpp
if(mp.count(x))
```

---

## Iterate

```cpp
for(auto it : mp)
{
    cout << it.first << " " << it.second;
}
```

---

# 25. Time Complexity Summary

| STL           | Data Structure | Search   | Insert   | Delete   |
| ------------- | -------------- | -------- | -------- | -------- |
| map           | Red Black Tree | O(log n) | O(log n) | O(log n) |
| unordered_map | Hash Table     | O(1) avg | O(1) avg | O(1) avg |

---

# 26. Interview-Level Insights

## Why BST instead of normal binary tree?

Normal BST can become skewed:

```text
1
 \
  2
   \
    3
```

Complexity becomes:

```text
O(n)
```

So `map` uses self-balancing tree.

---

## Why unordered_map Can Be Dangerous?

Some competitive programming hacks create many collisions intentionally.

Then:

```text
O(n²)
```

may happen.

---

# 27. Most Important HashMap Patterns in DSA

You should master these:

---

## A. Frequency Count

```cpp
freq[x]++;
```

---

## B. Visited Array Replacement

```cpp
unordered_map<int,bool> vis;
```

Useful when numbers are huge.

---

## C. Prefix Sum + HashMap

Very important for:

- subarray sum
- longest subarray
- zero sum problems

---

## D. Two Sum Problem

Classic interview problem.

```cpp
unordered_map<int,int> mp;
```

Stores:

```text
value -> index
```

---

# 28. Best Practice

## Use `unordered_map` when:

- speed needed
- ordering unnecessary

## Use `map` when:

- sorted order needed
- range queries needed

---

# 29. Final Mental Model

## map

```text
Balanced Tree
Sorted
Stable O(log n)
```

---

## unordered_map

```text
Hash Table
Very Fast
Average O(1)
Worst O(n)
```

---

# 30. Golden Interview Line

You can say:

> "`map` uses a self-balancing BST, so operations are O(log n). `unordered_map` uses hashing, giving average O(1) access, but collisions can degrade it to O(n)."

That single sentence already sounds interview-ready.
