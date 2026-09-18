# geomanji-backend

Go backend that geolocates an image via the geomanji API and returns the
result, optionally as a Google Maps or Google Earth link.

## Setup

```bash
go mod tidy
cp .env.example .env   # if you keep one; otherwise set env vars directly
go run main.go
```

### Environment variables

| Variable          | Required | Default | Description                          |
|-------------------|----------|---------|---------------------------------------|
| `PORT`            | No       | `8080`  | Port the HTTP server listens on       |
| `GEOMANJI_API_URL` | Yes     | -       | Base URL of the geomanji geolocate API |

No Google API key is needed — the Maps/Earth links use public URL schemes,
not the Maps Platform APIs.

## Endpoints

### `GET /geolocate`
Health/placeholder route.

### `POST /search`
Geolocates an uploaded image.

```bash
curl -X POST http://localhost:8080/search \
  -F "img=@sample.png"
```

```json
{
  "Latitude": 12.97,
  "Longitude": 77.59,
  "Probability": 0.83
}
```

### `POST /search/maps`
Same as `/search`, plus a Google Maps link pinning the location.

```json
{
  "Latitude": 12.97,
  "Longitude": 77.59,
  "Probability": 0.83,
  "MapsLink": "https://www.google.com/maps/search/?api=1&query=12.970000,77.590000"
}
```

### `POST /search/earth`
Same as `/search`, plus a Google Earth link with a tilted, zoomed-in 3D
camera; Earth animates the fly-in itself when the link is opened.

```json
{
  "Latitude": 12.97,
  "Longitude": 77.59,
  "Probability": 0.83,
  "EarthLink": "https://earth.google.com/web/@12.970000,77.590000,0a,300d,35y,45h,65t,0r"
}
```

## Project layout

```
main.go                              # entrypoint, env loading, server start
pkg/http/router.go                   # route table
pkg/http/server.go                   # http.Server construction
pkg/http/controllers/                # request handlers (DI via Handler struct)
pkg/integrations/geomanji.go         # geomanji API client
pkg/integrations/maps.go             # Google Maps / Earth link builders
```

Add new dependencies (DB, clients, config) as fields on `controllers.Handler`
and pass them in from `main.go`.
