const cacheName = "app-" + "ef2c64dacb82e213d5b7e08a15922c062e1bed20";

self.addEventListener("install", event => {
  console.log("installing app worker ef2c64dacb82e213d5b7e08a15922c062e1bed20");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/papilio",
          "/papilio/app.css",
          "/papilio/app.js",
          "/papilio/manifest.webmanifest",
          "/papilio/wasm_exec.js",
          "/papilio/web/app.wasm",
          "/papilio/web/default.png",
          "/papilio/web/large.png",
          "https://unpkg.com/@patternfly/patternfly@4.164.2/patternfly-addons.css",
          "https://unpkg.com/@patternfly/patternfly@4.164.2/patternfly.css",
          
        ]);
      }).
      then(() => {
        self.skipWaiting();
      })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker ef2c64dacb82e213d5b7e08a15922c062e1bed20 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
