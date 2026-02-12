# Library

## Directory Purpose

- `.storage/`: Local storage area for generated/runtime files (for example cache files, temporary exports, or files created by the app while running).
- `config/`: Application configuration package and configuration-loading logic.
- `migrations/`: Database migration scripts used to create/update database schema.
- `package/`: Shared reusable library code (helpers, models, utilities, logging, crypto, validation, pagination, reporting, wrappers, etc.).
- `webserver/controllers/`: HTTP handlers/controllers that receive requests and orchestrate application actions.
- `webserver/resources/`: API resource/response mappers used to shape data returned by endpoints.
- `webserver/middlewares/`: HTTP middleware components (auth, cors, logging, recovery, rate limit, security, token/session handling).
- `webserver/routes/`: Route registration and endpoint mapping to controllers.
- `services/entity/`: Domain entities and core business data structures.
- `services/database/`: Database connection setup and database service wiring.
- `services/error_message/`: Centralized error definitions/messages used across services.
- `services/repository/`: Data access layer for querying and persisting entities.
- `services/usecase/`: Business use cases/application service layer containing core business logic.

## Git Init Commands

```bash
echo "# library" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:yonakangwe/library.git
git push -u origin main
```
