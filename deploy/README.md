# Docker Compose: API + admin

This deployment intentionally serves only the API, uploads, health endpoint, and admin SPA.
The customer storefront is deployed separately (for example on Vercel).

## First deployment

On the server, copy this directory and create the two untracked deployment files:

```bash
cp .env.example .env
cp ../config.yml.example ./config.yml
```

Set `DJ_IMAGE` in `.env` to the published image, for example
`ghcr.io/<github-owner>/dujiao-faka:v1.2.3`. For a private GHCR package, authenticate once with
a GitHub personal access token that has `read:packages`:

```bash
echo "$GHCR_TOKEN" | docker login ghcr.io -u <github-user> --password-stdin
```

Edit `config.yml` before starting:

- Replace `app.secret_key`, `jwt.secret`, and `user_jwt.secret` with three different `openssl rand -hex 32` values.
- Set `server.mode: release`.
- Set `redis.host: redis` and `queue.host: redis`.
- Set `cors.allowed_origins` to your Vercel production URL, such as `https://shop.example.com`; do not use `*` in production.
- Set `web.admin_path` to an unguessable path, such as `/dj-mgmt-7x9k2`.
- Set `server.trusted_proxies` to the exact source IP/CIDR that the app sees for your Nginx/Caddy proxy. With a host proxy forwarded through Docker this is commonly the Docker bridge gateway, not necessarily `127.0.0.1`; inspect your Docker network and permit only that proxy address/range.

Start it:

```bash
docker compose pull
docker compose up -d
docker compose logs -f app
```

Point a TLS reverse proxy at `http://127.0.0.1:8080`. The admin panel is then available at
`https://api.example.com/<web.admin_path>` and the API at `https://api.example.com/api/v1`.

## Vercel storefront

Set the Vercel environment variable below at build time, then redeploy:

```text
VITE_API_BASE_URL=https://api.example.com
```

This value must exactly match an entry in `cors.allowed_origins` on the API server (the Vercel
origin, not the API origin). Also configure payment callbacks and OAuth/Telegram redirect URLs
to use the API domain.

## Update and rollback

Prefer an immutable release tag in `DJ_IMAGE`, then update it and run:

```bash
docker compose pull app
docker compose up -d
```

To roll back, change `DJ_IMAGE` back to the previous tag and run the same commands. Named
volumes retain SQLite data, uploads, logs, and Redis data across image updates.
