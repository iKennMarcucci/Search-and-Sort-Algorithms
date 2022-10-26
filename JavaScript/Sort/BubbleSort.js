import {printMsg1, printMsg2} from '../NumberGen/RandomInt.js'
// Bubble sort Implementation using Javascript

function bblSort(arr) {
  let it = 0;
  for (var i = 0; i < arr.length; i++) {
    // Last i elements are already in place
    for (var j = 0; j < (arr.length - i - 1); j++) {
      // Checking if the item at present iteration
      it++;
      // is greater than the next iteration
      if (arr[j] > arr[j + 1]) {
        // If the condition is true then swap them
        var temp = arr[j]
        arr[j] = arr[j + 1]
        arr[j + 1] = temp
      }
    }
  }
  process.stdout.write("\nBUBBLE SORT\n");
  process.stdout.write("Iteraciones = " + it + "\n");
  return arr;
}

//Unsorted array
//TEST
// let arr = [234, 43, 55, 63, 5, 6, 235, 547];

// printMsg1(arr);
// printMsg2(bblSort(arr));
// Now pass this array to the bblSort() function

export {bblSort as bubbleSort};