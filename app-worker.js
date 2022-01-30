const cacheName = "app-" + "ad869b856df65cf4eb9eaeba096f74a60e2d0790";

self.addEventListener("install", event => {
  console.log("installing app worker ad869b856df65cf4eb9eaeba096f74a60e2d0790");

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
  console.log("app worker ad869b856df65cf4eb9eaeba096f74a60e2d0790 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
