# Interface

## Overview
* Simple Pipeline Job (Auto and **Manual**)
  * Each pipe is docker container
* Gkin has pipeline execute role
  * There is a difference in job focus between docker compose and gkin
* User doing
  * Create (Dockerfile) ← This is not decision
  * Create Pipeline Config(.gkin.yaml)
  * Run Web UI

### Pipeline Config
* Step executes
  * Restart from failed step
* Recive parameter