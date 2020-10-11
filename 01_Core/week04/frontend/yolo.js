fetchTours();

function fetchTours() {

    fetch(`http://localhost:8000/frontend/yolo.html`)
    .then(response => {
        return response.json();
    })
    .then(data => {
        displayDom(data);
    })
}




function displayDom(d) {
    for (i = 0; i < d.length; i++) {
        document.getElementById("Name" + i.toString()).innerHTML = d[i].product_name;
        document.getElementById("Duration" + i.toString()).innerHTML = d[i].duration;
        document.getElementById("Cost" + i.toString()).innerHTML = d[i].cost;
        document.getElementById(i + 1).src = d[i].img;
    }
}