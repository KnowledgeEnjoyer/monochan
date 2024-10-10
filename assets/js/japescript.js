function id(x) {
    return document.getElementById(x)
}

function ids(x) {
    return document.querySelectorAll(`[id=${x}]`)
}

function classes(x) {
    return document.getElementsByClassName(x)
}

function tags(x) {
    return document.getElementsByTagName(x)
}

function parents(x) {
    let y = x
    z = []
    if (y != null) {
        while (y) {
            z.push(y);
            y = y.parentNode;
        }
        return z
    }
}

// HTMLElement.prototype.add = function(x) {
//     this.x.insertBefore(x, this.nextSibling);
// };

// function addBefore(referenceNode, newNode) {
//     referenceNode.parentNode.insertBefore(newNode, referenceNode.nextSibling);
// }

function create(x, y) {
    let z = document.createElement(x);
    z.appendChild(document.createTextNode(y));
    return z;
}

function print(x) {
    return console.log(x)
}