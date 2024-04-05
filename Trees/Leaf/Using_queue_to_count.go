package Leaf

/*
Step 1 − Create a package main and declare fmt(format package) and strings package in the program where
			main produces executable codes and fmt helps in formatting input and output.

Step 2 − Create a struct called TreeNode that has fields for the node value and pointers to
			its left and right children to represent a node in the tree.

Step 3 − Create a function called "count_leaf_nodes" that takes a TreeNode pointer as an argument and
			returns the number of leaf nodes in the tree.

Step 4 − Create a queue and set its root node and count to 0.

Step 5 −

a. Dequeue the initial node from the queue while it still contains nodes.

b. Increase the count if the dequeued node is a leaf node (that is, if its left and right children are nil).

c. Enqueue the left child of the dequeued node if it is not null.

d. Enqueue the right child of the dequeued node if it is not null.

Step 6 − Return the count to the function.

Step 7 − Create a TreeNode struct instance in the main function to represent the tree's root node, and set its left and right children to other TreeNode struct instances to represent the rest of the tree.
*/
