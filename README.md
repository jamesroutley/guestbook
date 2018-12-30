
```js
window.onload = function() {
  var xhr = new XMLHttpRequest();
  xhr.open("POST", "localhost:8080/log");
  xhr.send(
    JSON.stringify({
      url: window.location.href,
      referrer: document.referrer
    })
  );
};
```
