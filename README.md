# Guestbook

A simple web analytics server. I wrote it because I didn't want Google Analytics
on my websites but I wanted to know how many people were visiting certain pages.

## Server

The server exposes a single endpoint: 
```
/log {
    "url": "...",
    "referrer": "..."
}
```

The server stores this info, along with the user's IP address and a timestamp in
a SQLite database.

### Analytics

The server also exposes an analytics endpoint on port 4000. `curl`ing this
endpoint will return some statistics about recent visits.

## Client

Your website should run this javascript on every page:

```js
window.onload = function() {
  var xhr = new XMLHttpRequest();
  xhr.open("POST", "https://<your domain>/log");
  xhr.send(
    JSON.stringify({
      url: window.location.href,
      referrer: document.referrer
    })
  );
};
```

