# devcontainer.json

- **name** — Display name shown in VS Code's UI (bottom-left corner, container picker). *Not essential*, cosmetic only.
- **dockerComposeFile** — Tells VS Code to use Docker Compose instead of a single image/Dockerfile. The path is relative to devcontainer.json itself. *Essential* — this is what wires up the multi-container setup.
- **service** — Which service inside the compose file VS Code should attach to (open a terminal in, run extensions in, mount your files into). Must match a service name in docker-compose.yml. *Essential* — without it VS Code doesn't know which container is "yours".
- **workspaceFolder** — The path inside the container where VS Code opens. Must match where the volume is mounted in docker-compose.yml. *Essential* — if these disagree, VS Code opens to the wrong directory or fails to find your files.
- **shutdownAction** — What happens when you close the VS Code window. `stopCompose` stops all services in the compose file. The alternative is `none` (containers keep running). *Not essential*, but good hygiene — without it containers would keep running in the background after you close the window.
- **customizations.vscode.extensions** — 
Extensions VS Code automatically installs inside the container when it first builds. *Not strictly essential* (you could install it manually), but very convenient

# nginx/html/index.html

- **fetch(url)** — a built-in browser function that sends an HTTP GET request to the given URL. It doesn't return the response immediately; it returns a *Promise* (a placeholder for a value that will arrive later).
- **await** — pauses execution of this function *until the Promise resolves* (i.e., until the response arrives). The rest of the page keeps working normally; only this function is paused.

```js
// Fetch a Promise, but store it when its resolved to a Response
const res = await fetch('http://localhost:8090/hello');

// The headers arrive first, but the body is still streaming in.
// Because of this network design, reading the body is also asynchronous.
const data = await res.json();
```

# nginx/nginx.conf

- **events {}** — The events block configures how nginx handles connections at the *OS level* — things like how many simultaneous connections a worker can handle. The block is *required* by nginx or it refuses to start. Empty braces are valid and means "use defaults".
- **http {}** — Opens the *HTTP context* — everything related to serving web traffic goes inside. The alternative is a `stream {}` block for raw *TCP/UDP proxying*.
  - **include /etc/nginx/mime.types;** — Loads a file that *maps file extensions to Content-Type headers* — e.g. .html → text/html, .js → application/javascript. Without this, nginx doesn't know what Content-Type to send, so browsers may mishandle files (e.g. treat a .js file as plain text and refuse to execute it). The file ships with the nginx image so it's always there.
  - **default_type application/octet-stream;** — The fallback *Content-Type for any file extension not found in mime.types*. octet-stream means "generic binary" — browsers will prompt a download rather than try to render it. This is the safe default: unknown file = don't guess, force a download.
  Alternative: `text/plain` — would display unknown files as text in the browser. Less safe.
  - **server {}** — Defines a *virtual server* — one nginx instance can serve multiple domains/ports by having multiple server blocks.
    - **listen 80;** — *The port nginx listens on* inside the container. Port 80 is the standard HTTP port.
    - **root /usr/share/nginx/html;** — The filesystem path nginx uses as *the base when looking for files to serve*. A request for `/index.html` resolves to `/usr/share/nginx/html/index.html` on disk.
    - **index index.html;** — When a request comes in for a directory (e.g. `/`), *nginx serves this file automatically*. Without it, a request to `/` would either show a directory listing or a 403 — depending on whether autoindex is on.
    - **location / {}** — A *routing rule* that matches every request path (since `/` is a prefix, it catches everything). More specific location blocks take priority over less specific ones.
    - **try_files $uri $uri/ =404;** — For each request, *nginx tries in order*:
      1. `$uri` — look for an exact *file* match (e.g. /about.html → looks for that file)
      2. `$uri/` — look for a *directory* with an index.html inside it
      3. `=404` — if neither exists, return a 404
