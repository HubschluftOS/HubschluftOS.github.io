function myFunction() {
  var input, filter, ul, li, a, i;
  input = document.getElementById("mySearch");
  filter = input.value.toUpperCase();
  ul = document.getElementById("myMenu");
  li = ul.getElementsByTagName("li");

  for (i = 0; i < li.length; i++) {
    a = li[i].getElementsByTagName("a")[0];
    if (a.innerHTML.toUpperCase().indexOf(filter) > -1) {
      li[i].style.display = "";
    } else {
      li[i].style.display = "none";
    }
  }
}

function count() {
  var myMenu = document.getElementById("myMenu");
  var listItems = myMenu.getElementsByTagName("li");
  var countOutput = document.getElementById("countOutput");

  countOutput.textContent += listItems.length;
}

count();

function sort() {
  var myMenu = document.getElementById("myMenu");
  var listItems = myMenu.getElementsByTagName("li");
  var itemsArray = Array.from(listItems);
  itemsArray.sort(function (a, b) {
    var textA = a.textContent || a.innerText;
    var textB = b.textContent || b.innerText;
    return textA.localeCompare(textB);
  });

  myMenu.innerHTML = "";

  itemsArray.forEach(function (item) {
    myMenu.appendChild(item);
  });
}

sort();
