
function linearSearch(arr, key) {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === key) {
            return i
        }
    }
    return -1
}

// let arr = [10, 22, 35, 40, 45, 50, 80, 82, 85, 90, 100, 235];
// let x = 235;
// let ind = linearSearch(arr, x);


// if (ind >= 0) {
//     console.log("\nElement found " + (++ind));
// } else {
//     console.log(x + " Element not found!\n");
// }

export {linearSearch}