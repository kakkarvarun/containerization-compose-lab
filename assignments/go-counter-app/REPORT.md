# Short Report

**App & Steps:**  
I built a tiny Go web app that serves on port 8080, reads a greeting from `APP_GREETING`, and stores a request counter in `/app/data/counter.txt`. I created a multi-stage Dockerfile: Stage 1 uses `golang:1.22.5-alpine` to compile a static binary with an injected version; Stage 2 uses `alpine:3.20.3` as the small runtime. I created a non-root `app` user, set `EXPOSE 8080`, and declared `/app/data` as a volume.

**Build:**  
From `assignments/go-counter-app`, I ran:
docker build -t vk-go-counter:1.0.0 --build-arg VERSION=1.0.0 .
I pinned image versions (no `latest`) and used a semantic tag (`1.0.0`) for the output image.

**Run & Volumes/Bind Mounts:**  
I created a local `data/` folder and ran:
docker run --name go-counter -p 8085:8080 -e APP_GREETING="Hello Varun" -v "${PWD}\data:/app/data" -d vk-go-counter:1.0.0
The app writes the hit counter to `/app/data/counter.txt`. Because that path is bind-mounted to `./data` on the host, the counter persists across container restarts (`docker stop` / `docker start`) and removals (`docker rm`) as long as the host folder remains.

**Best Practices:**  
- Multi-stage build to keep the final image small.  
- Pinned base images for reproducibility.  
- Non-root user for better security.  
- Minimal layers and a `.dockerignore` to avoid copying unneeded files.  
- `EXPOSE 8080` documents the container port.  
- Versioned image tag (`vk-go-counter:1.0.0`) instead of `latest`.

**Challenges & Resolutions:**  
On Windows, path syntax for bind mounts can be tricky. I used PowerShell’s `${PWD}\data:/app/data` form to ensure the mount worked. I also verified logs with `docker logs go-counter` and confirmed persistence by refreshing the page and seeing the counter increase after restarts.
