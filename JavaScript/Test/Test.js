import { fillArr,fillArrmM, fillArrMm, fillArrPS, printArray } from '../NumberGen/RandomInt.js';

import {bubbleSort} from '../Sort/BubbleSort.js'
import {insertionSort} from '../Sort/InsertionSort.js'
import {quickSort} from '../Sort/QuickSort.js'
import {selectionSort} from '../Sort/SelectionSort.js'
import {shellSort} from '../Sort/ShellSort.js'

import {binarysearch} from '../Search/BinarySearch.js'
import {fibonacciSearch} from '../Search/FibonacciSearch.js'
import {linearSearch} from '../Search/LinearSearch.js'

//100, 1000, 1000000, 1000000000
let n = 1000;
// Reemplace fillArr => fillArrmM , fillArrMm, fillArrPS
let arr = fillArr(n);
// let x = arr[n/4];

//BINARY SEARCH
//BinarySearch && BubbleSort  
let arr1 = bubbleSort(arr);
binarysearch(arr1);
fibonacciSearch(arr1);
linearSearch(arr1)

//BinarySearch && Insertion
let arr2 = insertionSort(arr);
binarysearch(arr2);
fibonacciSearch(arr2);
linearSearch(arr2);

//BinarySearch && QuickSort
let arr3 = quickSort(arr);
binarysearch(arr3);
fibonacciSearch(arr3);
linearSearch(arr3);

//BubbleSort & Selection
let arr4 = selectionSort(arr);
binarysearch(arr4);
fibonacciSearch(arr4);
linearSearch(arr4);

//BinarySearch && ShellSort
let arr5 = shellSort(arr);
binarysearch(arr5);
fibonacciSearch(arr5);
linearSearch(arr5);

process.stdout.write("\n");
printArray(arr);