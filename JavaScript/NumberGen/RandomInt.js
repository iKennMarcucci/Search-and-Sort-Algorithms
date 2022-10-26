//Get a random 5 digits number
function getRandomInt() {
    let num = Math.floor(Math.random() * 90000) + 10000;
    return num;
}

//Generate a 5 digits array with n lenght

//Unsorted
function fillArr(n) {
    var arr = new Array(n);
    for (var i = 0; i < n; i++) {
        arr[i] = getRandomInt();
    }
    return arr;
}

//Sorter minor to major
function fillArrmM(n) {
    var arr = new Array(n);
    for (var i = 0; i < n; i++) {
        arr[i] = getRandomInt();
    }
    
    return arr.sort();
}
//Sorter major to minor
function fillArrMm(n) {
    var arr = new Array(n);
    var res = new Array(n);
    
    for (var i = 0; i < n; i++) {
        arr[i] = getRandomInt();
    }
    arr.sort();
    for (var i = 0, j = n-1; i < n; i++, j--) {
        res[i] = arr[j];
    }
    return res;
}
//Partial sorted

function fillArrPS(n) {
    var arr = new Array(n);
    var res = new Array(n);

    for (var i = 0; i < n/2; i++) {
        arr[i] = getRandomInt();
    }
    res = arr.sort();
    for (var i = n/2 ; i < n; i++) {
        res[i] =  getRandomInt();
    }
    return res;
}
// Print a array
function printArray(arr) {
    process.stdout.write(" [ ");
    let n = arr.length;
    for (let i = 0; i < n; i++) {
        process.stdout.write(arr[i] + " ");
    }
    process.stdout.write("]");
}

//Print iterations and sorted/unsorted
function printMsg1(arr) {
    process.stdout.write("\nUnsorted array:");
    printArray(arr);
}

function printMsg2(arr) {
    process.stdout.write("\nSorted array:");
    printArray(arr);
}

export {fillArr, printArray, printMsg1, printMsg2, fillArrmM, fillArrMm,fillArrPS};