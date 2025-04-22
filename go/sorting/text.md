**Sorting Algorithm Properties:**

* **Stable:** Preserves the relative order of equal elements.
    * Example: `[(Alice, 90), (Charlie, 90)]` will remain in the same order after sorting by grade.
* **In-Place:** Sorts elements within the original array, using minimal extra memory (O(1)).
    * Example: Bubble sort, insertion sort.
* **Adaptable:** Performs better (faster) when the input is partially sorted.
    * Example: Insertion sort.


| Algorithm       | Worst-Case Time | Average-Case Time | Best-Case Time | Space Complexity | Stable | In-Place | Adaptable | Description                                                                     |
|-----------------|-----------------|-------------------|----------------|------------------|--------|----------|-----------|---------------------------------------------------------------------------------|
| Bubble Sort     | O(n^2)          | O(n^2)            | O(n)           | O(1)             | Yes    | Yes      | Yes       | Repeatedly compares adjacent elements, swapping if out of order.                 |
| Insertion Sort  | O(n^2)          | O(n^2)            | O(n)           | O(1)             | Yes    | Yes      | Yes       | Builds the sorted array one element at a time, inserting into correct position. |
| Selection Sort  | O(n^2)          | O(n^2)            | O(n^2)         | O(1)             | No     | Yes      | No        | Repeatedly finds the minimum element and places it at the beginning.           |
| Merge Sort      | O(n log n)      | O(n log n)        | O(n log n)     | O(n)             | Yes    | No       | No        | Divides the array, recursively sorts halves, and merges them.                            |
| Quick Sort      | O(n^2)          | O(n log n)        | O(n log n)     | O(log n)         | No     | Yes      | No        | Selects a pivot, partitions elements, and recursively sorts sub-arrays.            |
| Heap Sort       | O(n log n)      | O(n log n)        | O(n log n)     | O(1)             | No     | Yes      | No        | Builds a heap and repeatedly extracts the maximum element.                               |
| Counting Sort   | O(n + k)        | O(n + k)          | O(n + k)       | O(k)             | Yes    | No       | No        | Counts occurrences of each element and uses counts to determine positions.              |
| Radix Sort      | O(nk)           | O(nk)             | O(nk)          | O(n + k)         | Yes    | No       | No        | Sorts elements by processing individual digits or characters.                                |
| Bucket Sort     | O(n^2)          | O(n + k)          | O(n)           | O(n + k)         | Yes    | No       | Yes       | Distributes elements into buckets, sorts buckets, and concatenates.                      |