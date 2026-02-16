# MWONGOZO WA MRADI WA LIBRARY - GOLANG

## MUHTASARI
Mradi huu ni REST API iliyojengwa kwa kutumia lugha ya Go (Golang) kwa ajili ya kusimamia mfumo wa maktaba. Inatumia muundo wa **Clean Architecture** wenye tabaka tofauti za Controller, UseCase, Repository, na Entity.

---

## MUUNDO WA MRADI

```
library/
├── app.go                          # Faili kuu ya kuanzisha programu
├── config/                         # Usanidi wa mfumo
│   └── config.go
├── config.yml                      # Faili ya usanidi
├── webserver/                      # Tabaka la webserver
│   ├── webserver.go               # Kuanzisha webserver
│   ├── routes/                    # Njia za API
│   │   ├── router.go
│   │   └── api.go
│   ├── controllers/               # Vidhibiti vya request
│   │   ├── role_controller.go
│   │   ├── mkoa_controller.go
│   │   └── staff_controller.go
│   └── middlewares/               # Middleware za usalama
│       ├── cors.go
│       ├── logger.go
│       ├── secure.go
│       └── recover.go
├── services/                      # Tabaka la biashara
│   ├── usecase/                   # Mantiki ya biashara
│   │   └── role/
│   │       ├── role_service.go
│   │       └── role_interface.go
│   ├── repository/                # Mawasiliano na database
│   │   └── role_repository.go
│   ├── entity/                    # Miundo ya data
│   │   └── role_entity.go
│   └── database/                  # Muunganisho wa database
│       └── database.go
└── package/                       # Vifaa vya msaada
    ├── models/                    # Miundo ya JSON
    ├── validator/                 # Uthibitishaji wa data
    ├── wrappers/                  # Majibu ya API
    └── pagination/                # Upangaji wa kurasa

```

---

## MTIRIRIKO WA DATA: KUSAJILI MFANYAKAZI MPYA (CREATE STAFF)

### **MUHTASARI WA MTIRIRIKO**

Mradi huu unatumia **Clean Architecture** ambapo mtiririko wa data unafuata mpangilio huu:

```
1. ENTITY (Muundo wa Data) → Inaelezea jinsi data inavyoonekana
2. USECASE/SERVICE (Mantiki ya Biashara) → Inafanya kazi za biashara
3. REPOSITORY (Mawasiliano na Database) → Inahifadhi data
4. MODEL (Muundo wa API) → Inaelezea JSON inayotumwa/kupokelewa
5. CONTROLLER (Kupokea Request) → Inapokea na kusafisha data
6. ROUTE (Njia za API) → Inaelekeza request kwenye controller sahihi
7. DATABASE (PostgreSQL) → Inahifadhi data kabisa
```

Hebu tuangalie kila tabaka kwa undani kwa kutumia mfano wa **kusajili mfanyakazi (Staff)**:

---

## TABAKA 1: ENTITY - MUUNDO WA DATA (`services/entity/staff_entity.go`)

**Entity** ni muundo wa data wa ndani wa mfumo. Inaelezea jinsi data inavyoonekana na ina sheria za uthibitishaji.

```go
// Mstari 11-25: Muundo wa Entity
type Staff struct {
    ID           int32         // Nambari ya utambulisho (PK)
    FullName     string        // Jina kamili la mfanyakazi
    Email        string        // Barua pepe
    Phone        string        // Nambari ya simu
    Username     string        // Jina la mtumiaji
    PasswordHash string        // Nywila iliyosimbwa (bcrypt hash)
    IsActive     bool          // Ameamilishwa?
    CreatedBy    int32         // Nani aliyemsajili
    CreatedAt    time.Time     // Wakati wa kusajili
    UpdatedBy    *int32        // 👈 Pointer: Nani alisasisha (NULL ikiwa bado)
    UpdatedAt    *time.Time    // 👈 Pointer: Wakati wa kusasisha (NULL ikiwa bado)
    DeletedBy    *int32        // 👈 Pointer: Nani alifuta (NULL ikiwa bado)
    DeletedAt    *time.Time    // 👈 Pointer: Wakati wa kufuta (NULL ikiwa bado)
}
```

**Maelezo:**
- Entity ni "blueprint" ya data - inaelezea sifa zote za Staff
- Ina fields za audit trail (CreatedBy, UpdatedBy, DeletedBy)
- Ina `PasswordHash` kwa usalama wa nywila
- **👈 Pointers (`*int32`, `*time.Time`):** Kwa ajili ya kushughulikia NULL values kutoka PostgreSQL
- Inatofautiana na Model (inayotumika kwa API)

```go
// Mstari 27-49: Constructor - Kusajili Staff mpya
func NewStaff(fullname, email, phone, username, passwordHash string, createdBy int32) (*Staff, error) {
    
    // HATUA 1.1: Tengeneza instance ya Staff
    staff := &Staff{
        FullName:     fullname,              // Weka jina kamili
        Email:        email,                 // Weka email
        Phone:        phone,                 // Weka simu
        Username:     username,              // Weka username
        PasswordHash: passwordHash,          // Weka nywila (bado haijasimbwa)
        CreatedBy:    createdBy,             // Weka mtumiaji aliyemsajili
    }
    
    // HATUA 1.2: Simba nywila kwa usalama
    if err := staff.EncryptPassword(); err != nil {
        log.Errorf("error encrypting password %v", err)
        return nil, err
    }
    
    // HATUA 1.3: Thibitisha data kabla ya kuendelea
    err := staff.ValidateCreate()
    if err != nil {
        log.Errorf("error validating Staff entity %v", err)
        return nil, err                      // Rudisha hitilafu
    }
    
    return staff, nil                        // Rudisha entity iliyothibitishwa
}
```

**Maelezo:**
- `NewStaff()` ni constructor - inatengeneza instance mpya ya Staff
- **HATUA 1.2** inasimba nywila kwa kutumia bcrypt (usalama mkubwa!)
- Inaitisha `ValidateCreate()` kuhakikisha data ni sahihi
- Inarudisha `*Staff` (pointer) au hitilafu

```go
// Mstari 51-71: Uthibitishaji wa Data
func (r *Staff) ValidateCreate() error {
    // HATUA 1.4: Thibitisha jina kamili
    if r.FullName == "" {
        return errors.New("error validating Staff entity, name field required")
    }
    
    // HATUA 1.5: Thibitisha email
    if r.Email == "" {
        return errors.New("error validating Staff entity, email field required")
    }
    
    // HATUA 1.6: Thibitisha simu
    if r.Phone == "" {
        return errors.New("error validating Staff entity, phone field required")
    }
    
    // HATUA 1.7: Thibitisha username
    if r.Username == "" {
        return errors.New("error validating Staff entity, username field required")
    }
    
    // HATUA 1.8: Thibitisha nywila
    if r.PasswordHash == "" {
        return errors.New("error validating Staff entity, password_hash field required")
    }
    
    // HATUA 1.9: Thibitisha mtumiaji aliyemsajili
    if r.CreatedBy <= 0 {
        return errors.New("error validating Staff entity, created_by field required")
    }
    
    return nil    // Data ni sahihi!
}
```

**Maelezo kwa Kina:**

**Line 51:** `func (r *Staff) ValidateCreate() error`
- Hii ni **Method** - function inayofanya kazi kwenye struct
- `(r *Staff)` - Receiver, inamaanisha method inafanya kazi kwenye Staff
- `error` - Inarudisha error kama kuna tatizo, au `nil` kama ni sahihi

**Line 52-70:** Business Logic Validation
- Kila field inathibitishwa kama haikuwa tupu (`""`)
- `CreatedBy` lazima iwe kubwa kuliko 0
- `errors.New()` - Tengeneza error message mpya

**Kwa Nini Tofauti na Model Validation?**
- **Entity Validation** - Sheria za biashara (business rules)
- **Model Validation** - Sheria za format ya data (JSON validation)

```go
// Mstari 95-108: Kusimba Nywila kwa Usalama
func (r *Staff) EncryptPassword() error {
    // HATUA 1: Angalia kama nywila ipo
    if r.PasswordHash == "" {
        return errors.New("password is required")
    }
    
    // HATUA 2: Simba nywila kwa kutumia bcrypt
    // bcrypt.DefaultCost = 10 (rounds za hashing)
    hashed, err := bcrypt.GenerateFromPassword(
        []byte(r.PasswordHash),  // Convert string kuwa []byte
        bcrypt.DefaultCost        // Cost factor (10)
    )
    if err != nil {
        return err   // Rudisha error kama hashing imeshindwa
    }
    
    // HATUA 3: Weka nywila iliyosimbwa
    r.PasswordHash = string(hashed)  // Convert []byte kuwa string
    
    return nil
}
```

**Maelezo kwa Kina:**

**Line 99:** `bcrypt.GenerateFromPassword()`
- Algorithm salama ya kusimba nywila
- `DefaultCost = 10` - rounds 10 za hashing (salama + haraka)
- Inarudisha `[]byte` - slice ya bytes

**Line 105:** Type Conversion
- `string(hashed)` - Convert `[]byte` kuwa `string`
- Sasa `PasswordHash` ina thamani kama: `$2a$10$N9qo8uLOickgx2ZMRZoMy.Mqr...`

**Usalama:**
- Bcrypt ni **one-way encryption** - haiwezi kurudishwa
- Hata kama database imeibiwa, nywila haiwezi kuonekana
- Inatumia **salt** - kila nywila ina hash tofauti hata kama ni sawa

**Mfano wa Matumizi:**
```go
// Kusajili mfanyakazi mpya
staff, err := entity.NewStaff(
    "John Doe",           // Jina kamili
    "john@example.com",   // Email
    "0712345678",         // Simu
    "johndoe",            // Username
    "SecurePass123",      // Nywila (itasimbwa automatic)
    1                     // ID ya mtumiaji aliyemsajili
)
if err != nil {
    // Hitilafu ya validation au encryption
    log.Error(err)
}
// Sasa staff iko tayari kutumika na nywila imeshasimbwa!
```

---

## TABAKA 2: USECASE/SERVICE - MANTIKI YA BIASHARA (`services/usecase/staff/staff_service.go`)

**UseCase/Service** ni tabaka la mantiki ya biashara. Linaelekeza mtiririko wa data na kufanya kazi za biashara.

```go
// Mstari 9-18: Muundo wa Service
type Service struct {
    repo Repository    // Repository ya kuhifadhi data
}

// Constructor - Kutengeneza instance ya Service
func NewService() UseCase {
    repo := repository.NewStaff()    // Tengeneza repository ya staff
    return &Service{
        repo: repo,                  // Weka repository
    }
}
```

**Maelezo kwa Kina:**

**Line 9:** `type Service struct`
- Struct inayowakilisha service layer
- Ina field moja: `repo` ya aina `Repository` (interface)

**Line 13:** `func NewService() UseCase`
- Factory function ya kutengeneza service
- Inarudisha `UseCase` (interface) - si `*Service`
- Hii ni **Dependency Injection Pattern**

**Line 14:** `repository.NewStaff()`
- Tengeneza repository instance
- Repository inajua jinsi ya kuongea na database

**Line 16-17:** Dependency Injection
- `&Service{repo: repo}` - Weka repository ndani ya service
- Service sasa inaweza kuita methods za repository

**Faida za Dependency Injection:**
- ✅ **Testability** - Rahisi kuingiza mock repository kwa testing
- ✅ **Flexibility** - Unaweza kubadilisha database bila kubadilisha service
- ✅ **Separation of Concerns** - Service haijui chochote kuhusu database

```go
// Mstari 20-33: Funguo ya Kusajili Staff
func (s *Service) CreateStaff(fullname, email, phone, username, passwordHash string, createdBy int32) (int32, error) {
    
    // HATUA 2.1: Tengeneza Entity na thibitisha
    staff, err := entity.NewStaff(fullname, email, phone, username, passwordHash, createdBy)
    if err != nil {
        return 0, err    // Validation au encryption imeshindwa
    }
    
    // HATUA 2.2: Hifadhi kwenye database kupitia Repository
    staffID, err := s.repo.Create(staff)
    if err != nil {
        return 0, err    // Kuna hitilafu ya database
    }
    
    // HATUA 2.3: Rudisha ID ya staff aliyesajiliwa
    return staffID, nil
}
```

**Maelezo kwa Kina:**

**Line 20:** Method Signature
- `func (s *Service)` - Method inayofanya kazi kwenye Service struct
- Parameters 6 za kuingiza data
- `(int32, error)` - Rudisha ID au error (Go pattern ya error handling)

**Line 22:** Entity Creation
- `entity.NewStaff(...)` - Tengeneza entity na uihakikishe
- Entity inajithibitisha na kujisimbia nywila
- Kama kuna hitilafu, rudisha `0, err`
- **Muhimu:** Sio `staff.ID` kwa sababu staff inaweza kuwa nil kama validation imeshindwa

**Line 27:** Repository Call
- `s.repo.Create(staff)` - Pita entity kwenye repository
- Repository inahifadhi kwenye database
- `s.repo` - Repository iliyopewa kwa njia ya Dependency Injection

**Line 28:** Error Handling
- `return 0, err` - Rudisha 0 (sio valid ID) na error
- **Muhimu:** Tunarudisha `0` badala ya `staff.ID` kwa sababu kama kuna error, `staff` inaweza kuwa nil au isiosahihi
- Hii inazuia **nil pointer dereference panic**

**Mtiririko wa Data:**
```
Service.CreateStaff()
    ↓
entity.NewStaff() → Validation + Encryption
    ↓
s.repo.Create() → SQL INSERT
    ↓
Database → Store data
    ↓
Return staffID
```

**HATUA 2.3 - Kurudisha Matokeo:**
- Inarudisha ID ya staff aliyesajiliwa
- ID hii inatoka database (auto-generated)

**Kazi Nyingine za Service (Mifano):**
```go
// Mstari 35-42: Funguo ya Kusasisha Staff
func (s *Service) UpdateStaff(e *entity.Staff) (int32, error) {
    // HATUA 1: Thibitisha entity kabla ya kusasisha
    err := e.ValidateUpdate()
    if err != nil {
        return e.ID, err    // Rudisha ID na error
    }
    
    // HATUA 2: Ita repository kufanya update
    return s.repo.Update(e)   // Repository inafanya SQL UPDATE
}

// Mstari 44-46: Funguo ya Kufuta Staff
func (s *Service) DeleteStaff(e *entity.Staff) (int32, error) {
    return s.repo.Delete(e)   // Repository inafanya SQL DELETE
}

// Mstari 48-50: Funguo ya Kupata Staff
func (s *Service) GetStaff(id int32) (*entity.Staff, error) {
    return s.repo.Get(id)   // Repository inafanya SQL SELECT
}
```

**Maelezo:**
- `UpdateStaff()` - Inathibitisha kabla ya kusasisha, kisha inaita repo.Update()
- `DeleteStaff()` - Simple pass-through kwa repository
- `GetStaff()` - Rudisha `*entity.Staff` au error

**Faida za Service Layer:**
- Unaweza kuongeza mantiki ya biashara hapa (kama kutuma email ya welcome)
- Unaweza kuthibitisha sheria za biashara kabla ya kuhifadhi
- Unaweza kuunganisha services nyingine (kama audit trail)
- Unaweza kuongeza logic ya kuthibitisha kama email tayari imetumika

---

## TABAKA 3: REPOSITORY - MAWASILIANO NA DATABASE (`services/repository/staff_repository.go`)

**Repository** ni tabaka linaloshughulikia mawasiliano na database. Linatengeneza SQL queries na kutekeleza.

```go
// Mstari 18-30: Muundo wa Repository
type StaffConn struct {
    conn *pgxpool.Pool    // Muunganisho wa database
}

// Constructor - Kutengeneza instance ya Repository
func NewStaff() *StaffConn {
    conn, err := database.Connect()    // Pata muunganisho wa database
    if util.IsError(err) {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
    }
    return &StaffConn{
        conn: conn,    // Weka muunganisho
    }
}
```

**Maelezo kwa Kina:**

**Line 400:** `type StaffConn struct`
- Struct inayowakilisha connection ya database
- `pgxpool.Pool` - Connection pool ya PostgreSQL (thread-safe)
- Connection pool inaruhusu multiple requests kwa wakati mmoja

**Line 405:** `func NewStaff() *StaffConn`
- Factory function ya kutengeneza repository
- `database.Connect()` - Singleton pattern ya connection
- `pgxpool` - Driver ya Go kwa PostgreSQL

**Line 406-409:** Error Handling
- `database.Connect()` - Pata connection au error
- `util.IsError(err)` - Check kama kuna error
- `fmt.Fprintf(...)` - Print error kwenye stderr

**Connection Pool Benefits:**
- ✅ **Performance** - Reuse connections badala ya kufungua mpya kila wakati
- ✅ **Thread-Safe** - Multiple goroutines zinaweza kutumia pool
- ✅ **Resource Management** - Kikomo cha connections zilizofunguliwa

```go
// Mstari 32: Jina la jedwali
var StaffTableName string = "staff"

// Mstari 34-49: Query ya msingi ya kupata data
func getStaffQuery() string {
    return `SELECT 
                 id,
                 full_name,
                 email,
                 phone,
                 username,
                 password_hash,
                 created_by, 
                 created_at, 
                 updated_by, 
                 updated_at, 
                 deleted_by, 
                 deleted_at 
             FROM ` + StaffTableName
}
```

**Maelezo kwa Kina:**

**Line 422:** `var StaffTableName string = "staff"`
- Variable ya jina la jedwali
- Rahisi kubadilisha kama jina la jedwali likibadilika
- Tunatumia variable badala ya string literal

**Line 425:** `func getStaffQuery() string`
- Helper function ya kutengeneza SELECT statement
- Inarudisha SQL query kama string

**Line 428:** `full_name`
- **Muhimu:** PostgreSQL inatumia **snake_case** (underscore)
- Si camelCase kama Go (fullName)
- Hii ilikuwa chanzo cha error: `column "fullname" does not exist`

**Line 439:** `FROM ` + StaffTableName
- Concatenate table name kwenye query
- `StaffTableName` - Variable ya jina la jedwali

```go
// Mstari 51-63: Funguo ya Kuhifadhi kwenye Database
func (con *StaffConn) Create(e *entity.Staff) (int32, error) {
    var StaffID int32    // Variable ya ID itakayorudishwa
    
    // HATUA 3.1: Tengeneza SQL INSERT query
    query := `INSERT INTO staff 
              (full_name, email, phone, username, password_hash, created_by, created_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) 
              RETURNING id`
    
    // HATUA 3.2: Tekeleza query na pokea ID
    err := con.conn.QueryRow(
        context.Background(),    // Context ya Go
        query,                   // SQL query
        e.FullName,              // $1 - Jina kamili
        e.Email,                 // $2 - Email
        e.Phone,                 // $3 - Simu
        e.Username,              // $4 - Username
        e.PasswordHash,          // $5 - Nywila iliyosimbwa (hash)
        e.CreatedBy,             // $6 - ID ya mtumiaji aliyemsajili
        time.Now(),              // $7 - Wakati wa sasa
    ).Scan(&StaffID)             // Pokea ID iliyorudishwa
    
    // HATUA 3.3: Shughulikia hitilafu
    if util.IsError(err) {
        if err.Error() == error_message.ErrNoResultSet.Error() {
            return StaffID, nil
        }
        log.Errorf("error creating staff from table %v: %v", StaffTableName, err)
    }
    
    return StaffID, err    // Rudisha ID ya staff mpya
}
```

**Maelezo kwa Kina:**

**Line 450:** `func (con *StaffConn) Create(e *entity.Staff) (int32, error)`
- Method inayofanya kazi kwenye StaffConn
- Inapokea `*entity.Staff` (pointer)
- Inarudisha `staffID` au `error`

**Line 454-457:** SQL INSERT
- `INSERT INTO staff` - Jedwali la staff
- `VALUES ($1, $2, ...)` - **Prepared statement** (usalama dhidi ya SQL injection)
- `RETURNING id` - PostgreSQL inarudisha ID iliyotengenezwa

**Prepared Statements:**
- ✅ **Usalama** - Zuia SQL injection attacks
- ✅ **Performance** - Query inaweza kucachewa
- ✅ **Type Safety** - Database inathibitisha types

**Line 460-469:** Parameter Substitution
- `$1, $2, ...` - Placeholders za PostgreSQL
- `e.FullName, e.Email, ...` - Actual values
- `time.Now()` - Timestamp ya sasa

**Line 470:** `.Scan(&StaffID)`
- `QueryRow()` - Execute query inayorudisha row moja
- `.Scan(&StaffID)` - Pokea ID iliyorudishwa
- `&StaffID` - Pointer kwa ajili ya kuhifadhi thamani

**Line 472-476:** Error Handling
- `util.IsError(err)` - Check kama kuna error
- `error_message.ErrNoResultSet` - Check kama hakuna results
- `log.Errorf(...)` - Log error kwa ajili ya debugging

**Usalama:**
- **Prepared Statements** zinazuia SQL Injection
- Parameters zinasafishwa na driver ya database
- **Nywila imeshasimbwa** kabla ya kufika hapa (Entity layer)
- Hakuna hatari ya malicious SQL code
- Nywila ya wazi haiwezi kuonekana kwenye database

**Kazi Nyingine za Repository:**
```go
// Kupata orodha ya wafanyakazi
func (con *StaffConn) List(filter *entity.StaffFilter) ([]*entity.Staff, int32, error) {
    // Tengeneza query na filters
    // Tekeleza SELECT statement
    // Rudisha orodha ya staff
}

// Kupata mfanyakazi mmoja
func (con *StaffConn) Get(id int32) (*entity.Staff, error) {
    // SELECT * FROM staff WHERE id = $1
}

// Kusasisha taarifa za mfanyakazi
func (con *StaffConn) Update(e *entity.Staff) error {
    // UPDATE staff SET ... WHERE id = $1
}

// Kufuta kwa soft delete
func (con *StaffConn) SoftDelete(id, deletedBy int32) error {
    // UPDATE staff SET deleted_at = NOW() WHERE id = $1
}

// Kufuta kabisa
func (con *StaffConn) HardDelete(id int32) error {
    // DELETE FROM staff WHERE id = $1
}
```

---

## TABAKA 4: MODEL - MUUNDO WA API (`package/models/staff.go`)

**Model** ni muundo wa data unaotumiwa kwa API (JSON). Inatofautiana na Entity.

```go
// Mstari 8-21: Muundo wa Model
type Staff struct {
    ID           int32     `json:"id" validate:"numeric,required"`
    FullName     string    `json:"fullName" validate:"required"`
    Email        string    `json:"email" validate:"required"`
    Phone        string    `json:"phone" validate:"required"`
    Username     string    `json:"username" validate:"required"`
    PasswordHash string    `json:"passwordHash" validate:"required"`
    CreatedBy    int32     `json:"created_by" validate:"numeric,required"`
    UpdatedBy    int32     `json:"updated_by" validate:"numeric,required"`
    DeletedBy    int32     `json:"deleted_by" validate:"numeric,required"`
    UpdatedAt    time.Time `json:"updated_at"`
    CreatedAt    time.Time `json:"created_at"`
    DeletedAt    time.Time `json:"deleted_at"`
}
```

**Maelezo:**
- **JSON Tags** (`json:"fullName"`) - Inaelezea jinsi field inavyoonekana kwenye JSON
- **Validation Tags** (`validate:"required"`) - Sheria za validation
- Model hii inatumika kubind JSON kutoka request
- `PasswordHash` inakubali nywila ya wazi kutoka request, lakini itasimbwa ndani ya Entity

**Tofauti kati ya Model na Entity:**

| Model | Entity |
|-------|--------|
| Inatumika kwa API (JSON) | Inatumika ndani ya mfumo |
| Ina JSON tags | Haina JSON tags |
| Validation ya input | Validation ya business logic |
| Inabadilika kulingana na API | Inabadilika kulingana na database |
| Inakubali nywila ya wazi | Inasimba nywila |

**Mfano wa JSON (Request):**
```json
{
  "fullName": "John Doe",
  "email": "john@example.com",
  "phone": "0712345678",
  "username": "johndoe",
  "passwordHash": "SecurePass123"
}
```

**Kumbuka:** Ingawa field inaitwa `passwordHash`, mtumiaji anatuma nywila ya wazi. Jina hili linatumika kwa consistency na database, lakini nywila itasimbwa automatic ndani ya Entity layer.

**Mfano wa JSON (Response):**
```json
{
  "id": 25,
  "fullName": "John Doe",
  "email": "john@example.com",
  "phone": "0712345678",
  "username": "johndoe",
  "created_by": 1,
  "created_at": "2024-02-15T22:00:00Z"
}
```

**Kumbuka:** Response hairudishi `passwordHash` kwa usalama!

---

## TABAKA 5: CONTROLLER - KUPOKEA REQUEST (`webserver/controllers/role_controller.go`)

**Controller** ni tabaka linalopokea HTTP requests, kusafisha data, na kuita service.

```go
// Mstari 106-127: Controller ya Kutengeneza Role
func CreateRole(c echo.Context) error {
    
    // ============================================
    // HATUA 5.1: BIND JSON DATA KUTOKA REQUEST
    // ============================================
    m := &models.Role{}    // Tengeneza struct tupu
    
    if err := c.Bind(m); util.IsError(err) {
        log.Errorf("error binding Role : %v", err)
        return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
    }
    
    // ============================================
    // HATUA 5.2: WEKA MTUMIAJI ALIYETENGENEZA
    // ============================================
    m.CreatedBy = 1    // ID ya mtumiaji (kwa sasa hardcoded)
    
    // ============================================
    // HATUA 5.3: THIBITISHA DATA (VALIDATION)
    // ============================================
    customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)
    if err := customValidator.ValidateStructPartial(m, "Name", "Description", "CreatedBy"); err != nil {
        log.Errorf("error validating Role  model: %v", err)
        return wrappers.ValidationErrorResponse(c, err)
    }
    
    // ============================================
    // HATUA 5.4: ITA SERVICE LAYER
    // ============================================
    service := role.NewService()    // Tengeneza instance ya service
    _, err := service.Create(m.Name, m.Description, m.CreatedBy)
    if util.IsError(err) {
        log.Errorf("error creating new %v: %v", m.Name, err)
        return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
    }
    
    // ============================================
    // HATUA 5.5: RUDISHA JIBU LA MAFANIKIO
    // ============================================
    return wrappers.MessageResponse(c, http.StatusCreated, m.Name+" created successfully")
}
```

**Maelezo kwa Kina:**

**HATUA 5.1 - Binding Data:**
```go
m := &models.Role{}
c.Bind(m)
```
- `c.Bind(m)` inabadilisha JSON kutoka request body kuwa Go struct
- Echo framework inafanya hii kwa automatic
- Kwa mfano, JSON hii:
  ```json
  {
    "name": "Admin",
    "description": "Administrator role"
  }
  ```
  Inabadilishwa kuwa:
  ```go
  models.Role{
    Name: "Admin",
    Description: "Administrator role"
  }
  ```

**HATUA 5.2 - CreatedBy:**
```go
m.CreatedBy = 1
```
- Inaweka ID ya mtumiaji aliyetengeneza
- Kwa sasa ni hardcoded (1)
- Kawaida inapata kutoka JWT token ya authentication

**HATUA 5.3 - Validation:**
```go
customValidator.ValidateStructPartial(m, "Name", "Description", "CreatedBy")
```
- Inathibitisha fields zilizotajwa tu
- Inatumia validation tags kutoka Model:
  - `validate:"required"` - Field lazima iwe na value
  - `validate:"numeric"` - Lazima iwe nambari
- Kama validation imeshindwa, inarudisha 400 Bad Request

**HATUA 5.4 - Service Call:**
```go
service := role.NewService()
service.Create(m.Name, m.Description, m.CreatedBy)
```
- Inatengeneza instance ya service
- Inapitisha data kwenye service layer
- Service inashughulikia mantiki ya biashara

**HATUA 5.5 - Response:**
```go
wrappers.MessageResponse(c, http.StatusCreated, "Admin created successfully")
```
- Inarudisha jibu la mafanikio
- Status code: 201 Created
- Message: "Admin created successfully"

**Controllers Nyingine:**
```go
// Orodha ya roles
func ListRole(c echo.Context) error {
    // Bind filter → Validate → Service.List() → Response
}

// Pata role moja
func GetRole(c echo.Context) error {
    // Bind ID → Validate → Service.Get() → Response
}

// Sasisha role
func UpdateRole(c echo.Context) error {
    // Bind data → Validate → Service.Update() → Response
}

// Futa role (soft delete)
func SoftDeleteRole(c echo.Context) error {
    // Bind ID → Validate → Service.SoftDelete() → Response
}

// Futa kabisa (hard delete)
func DestroyRole(c echo.Context) error {
    // Bind ID → Validate → Service.HardDelete() → Response
}
```

---

## TABAKA 6: ROUTE - NJIA ZA API (`webserver/routes/api.go`)

**Route** inaelekeza HTTP requests kwenye controller sahihi.

```go
// Mstari 9-38: Kusajili Endpoints za API
func ApiRouters(app *echo.Echo) {
    
    // ============================================
    // HATUA 6.1: TENGENEZA KUNDI LA API
    // ============================================
    api := app.Group("library/api/v1")
    
    // ============================================
    // HATUA 6.2: SAJILI NJIA ZA ROLE
    // ============================================
    role := api.Group("/role")
    {
        role.POST("/list", controllers.ListRole)         // Orodha ya roles
        role.POST("/show", controllers.GetRole)          // Pata role moja
        role.POST("/create", controllers.CreateRole)     // Tengeneza role mpya ← HAPA!
        role.POST("/update", controllers.UpdateRole)     // Sasisha role
        role.POST("/delete", controllers.SoftDeleteRole) // Futa kwa soft delete
        role.POST("/destroy", controllers.DestroyRole)   // Futa kabisa
    }
    
    // ============================================
    // HATUA 6.3: SAJILI NJIA ZA STAFF
    // ============================================
    staff := api.Group("/staff")
    {
        staff.POST("/create", controllers.CreateStaff)
    }
}
```

**Maelezo:**

**HATUA 6.1 - API Group:**
```go
api := app.Group("library/api/v1")
```
- Inatengeneza prefix ya njia zote: `library/api/v1`
- Njia zote zitaanza na prefix hii

**HATUA 6.2 - Role Routes:**
```go
role := api.Group("/role")
role.POST("/create", controllers.CreateRole)
```
- Inatengeneza sub-group: `/role`
- Njia kamili: `POST library/api/v1/role/create`
- Inaelekeza request kwenye `controllers.CreateRole`

**URL Kamili:**
```
http://127.0.0.1:4601/library/api/v1/role/create
```

**Middleware za Route (`webserver/routes/router.go`):**
```go
// Mstari 14-34: Middleware za Jumla
func Routers(app *echo.Echo) {
    
    // Middleware zinazofanya kazi kabla ya request kufika controller
    app.Use(middlewares.Cors())       // Ruhusu CORS (Cross-Origin requests)
    app.Use(middlewares.Gzip())       // Compress majibu
    app.Use(middlewares.Logger(true)) // Rekodi kila request
    app.Use(middlewares.Secure())     // Ongeza security headers
    app.Use(middlewares.Recover())    // Zuia server kuzimika kwa panic
    
    // Weka validator
    app.Validator = validator.GetValidator()
    
    // Sajili njia za API
    ApiRouters(app)
}
```

**Maelezo ya Middleware:**
- **CORS**: Inaruhusu maombi kutoka domain nyingine
- **Gzip**: Inapunguza ukubwa wa data inayotumwa
- **Logger**: Inarekodi kila request (wakati, njia, status code)
- **Secure**: Inaongeza HTTP headers za usalama
- **Recover**: Inazuia server kuzimika kama kuna panic

---

## TABAKA 7: DATABASE - POSTGRESQL

**Database** ni mahali ambapo data inahifadhiwa kabisa.

### **Muundo wa Jedwali la Roles:**

```sql
CREATE TABLE roles (
    id           SERIAL PRIMARY KEY,           -- ID (auto-increment)
    name         VARCHAR(255) NOT NULL,        -- Jina la role
    description  TEXT NOT NULL,                -- Maelezo
    created_by   INTEGER NOT NULL,             -- Nani aliyetengeneza
    created_at   TIMESTAMP NOT NULL,           -- Wakati wa kutengeneza
    updated_by   INTEGER DEFAULT 0,            -- Nani alisasisha
    updated_at   TIMESTAMP,                    -- Wakati wa kusasisha
    deleted_by   INTEGER DEFAULT 0,            -- Nani alifuta
    deleted_at   TIMESTAMP                     -- Wakati wa kufuta (soft delete)
);
```

**Maelezo:**
- `SERIAL PRIMARY KEY` - ID inazalishwa automatic
- `NOT NULL` - Field lazima iwe na value
- `TIMESTAMP` - Inahifadhi tarehe na wakati
- `deleted_at` - Kwa soft delete (NULL = haijafutwa)

### **SQL Query ya INSERT:**

```sql
INSERT INTO roles (name, description, created_by, created_at) 
VALUES ('Admin', 'Administrator role with full access', 1, '2024-02-15 22:00:00') 
RETURNING id;
```

**Matokeo:**
```
id
---
15
```

### **Muunganisho wa Database (`services/database/database.go`):**

```go
// Mstari 12-29: Kuunganisha na Database
var once sync.Once              // Hakikisha inaunganishwa mara moja tu
var instance *pgxpool.Pool      // Instance ya connection pool
var err error

func Connect() (*pgxpool.Pool, error) {
    once.Do(func() {
        // Pata connection string kutoka config
        connectionString := config.GetDatabaseConnection()
        
        // Unganisha na PostgreSQL
        instance, err = pgxpool.Connect(context.Background(), connectionString)
        if err != nil {
            log.Errorf("unable to create a database instance")
        }
    })
    
    if err != nil {
        log.Errorf("unable to connect to database: %v\n", err)
        return nil, err
    }
    
    return instance, err
}
```

**Maelezo:**
- `sync.Once` - Inaunganisha mara moja tu (Singleton pattern)
- `pgxpool.Pool` - Connection pool ya PostgreSQL
- Connection string inatoka `config.yml`:
  ```
  host=127.0.0.1 dbname=library user=root password=Emmanuel@100 port=5432
  ```

---

## MTIRIRIKO KAMILI: KUTENGENEZA ROLE MPYA

### **REQUEST kutoka Mtumiaji:**

```http
POST http://127.0.0.1:4601/library/api/v1/role/create
Content-Type: application/json

{
  "name": "Admin",
  "description": "Administrator role with full access"
}
```

### **SAFARI YA DATA:**

```
┌─────────────────────────────────────────────────────────────────┐
│ 1. HTTP REQUEST                                                  │
│    POST /library/api/v1/role/create                             │
│    Body: {"name": "Admin", "description": "Administrator..."}   │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ↓
┌─────────────────────────────────────────────────────────────────┐
│ 2. MIDDLEWARE LAYER                                             │
│    [CORS] → Thibitisha origin                                   │
│    [Logger] → Rekodi: "POST /role/create - 127.0.0.1"         │
│    [Secure] → Ongeza security headers                           │
│    [Recover] → Weka panic handler                               │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ↓
┌─────────────────────────────────────────────────────────────────┐
│ 3. ROUTE LAYER (webserver/routes/api.go)                       │
│    Endpoint: POST /library/api/v1/role/create                   │
│    Controller: controllers.CreateRole                           │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ↓
┌─────────────────────────────────────────────────────────────────┐
│ 4. CONTROLLER LAYER (controllers/role_controller.go:106)       │
│                                                                  │
│    HATUA 4.1: Bind JSON → models.Role                          │
│    ┌──────────────────────────────────────┐                    │
│    │ m := &models.Role{}                  │                    │
│    │ c.Bind(m)                            │                    │
│    │ Result: m.Name = "Admin"             │                    │
│    │         m.Description = "Admin..."    │                    │
│    └──────────────────────────────────────┘                    │
│                                                                  │
│    HATUA 4.2: Weka CreatedBy = 1                               │
│                                                                  │
│    HATUA 4.3: Validate                                          │
│    ┌──────────────────────────────────────┐                    │
│    │ ValidateStructPartial(m,             │                    │
│    │   "Name",        ✓ = "Admin"         │                    │
│    │   "Description", ✓ = "Administrator" │                    │
│    │   "CreatedBy")   ✓ = 1               │                    │
│    └──────────────────────────────────────┘                    │
│                                                                  │
│    HATUA 4.4: Ita Service                                       │
│    ┌──────────────────────────────────────┐                    │
│    │ service := role.NewService()         │                    │
│    │ service.Create("Admin", "Admin...", 1)│                    │
│    └──────────────────────────────────────┘                    │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ↓
┌─────────────────────────────────────────────────────────────────┐
│ 5. SERVICE/USECASE LAYER (usecase/role/role_service.go:20)     │
│                                                                  │
│    HATUA 5.1: Tengeneza Entity                                  │
│    ┌──────────────────────────────────────┐                    │
│    │ role, err := entity.NewRole(         │                    │
│    │   "Admin",                            │                    │
│    │   "Administrator role...",            │                    │
│    │   1                                   │                    │
│    │ )                                     │                    │
│    └──────────────────────────────────────┘                    │
│                         │                                        │
│                         ↓                                        │
└─────────────────────────┼────────────────────────────────────────┘
                          │
                          ↓
┌─────────────────────────────────────────────────────────────────┐
│ 6. ENTITY LAYER (entity/role_entity.go:21)                     │
│                                                                  │
│    HATUA 6.1: Tengeneza struct                                  │
│    ┌──────────────────────────────────────┐                    │
│    │ role := &Role{                       │                    │
│    │   Name: "Admin",                     │                    │
│    │   Description: "Administrator...",   │                    │
│    │   CreatedBy: 1                       │                    │
│    │ }                                    │                    │
│    └──────────────────────────────────────┘                    │
│                                                                  │
│    HATUA 6.2: Validate Business Rules                           │
│    ┌──────────────────────────────────────┐                    │
│    │ ValidateCreate():                    │                    │
│    │   ✓ Name != ""                       │                    │
│    │   ✓ Description != ""                │                    │
│    │   ✓ CreatedBy > 0                    │                    │
│    └──────────────────────────────────────┘                    │
│                                                                  │
│    HATUA 6.3: Rudisha entity                                    │
│    ┌──────────────────────────────────────┐                    │
│    │ return role, nil                     │                    │
│    └──────────────────────────────────────┘                    │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ↓
┌─────────────────────────────────────────────────────────────────┐
│ 7. KURUDI KWA SERVICE (usecase/role/role_service.go:25)        │
│                                                                  │
│    HATUA 7.1: Ita Repository                                    │
│    ┌──────────────────────────────────────┐                    │
│    │ roleID, err := s.repo.Create(role)   │                    │
│    └──────────────────────────────────────┘                    │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ↓
┌─────────────────────────────────────────────────────────────────┐
│ 8. REPOSITORY LAYER (repository/role_repository.go:50)         │
│                                                                  │
│    HATUA 8.1: Pata database connection                          │
│    ┌──────────────────────────────────────┐                    │
│    │ conn := database.Connect()           │                    │
│    └──────────────────────────────────────┘                    │
│                                                                  │
│    HATUA 8.2: Tengeneza SQL query                               │
│    ┌──────────────────────────────────────┐                    │
│    │ query := `                           │                    │
│    │   INSERT INTO roles                  │                    │
│    │   (name, description,                │                    │
│    │    created_by, created_at)           │                    │
│    │   VALUES ($1, $2, $3, $4)            │                    │
│    │   RETURNING id                       │                    │
│    │ `                                    │                    │
│    └──────────────────────────────────────┘                    │
│                                                                  │
│    HATUA 8.3: Tekeleza query                                    │
│    ┌──────────────────────────────────────┐                    │
│    │ conn.QueryRow(                       │                    │
│    │   query,                             │                    │
│    │   "Admin",           // $1           │                    │
│    │   "Administrator...", // $2           │                    │
│    │   1,                 // $3           │                    │
│    │   time.Now()         // $4           │                    │
│    │ ).Scan(&RoleID)                      │                    │
│    └──────────────────────────────────────┘                    │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ↓
┌─────────────────────────────────────────────────────────────────┐
│ 9. DATABASE LAYER (PostgreSQL)                                  │
│                                                                  │
│    HATUA 9.1: Tekeleza INSERT                                   │
│    ┌──────────────────────────────────────┐                    │
│    │ INSERT INTO roles                    │                    │
│    │ (name, description, created_by,      │                    │
│    │  created_at)                         │                    │
│    │ VALUES                               │                    │
│    │ ('Admin',                            │                    │
│    │  'Administrator role...',            │                    │
│    │  1,                                  │                    │
│    │  '2024-02-15 22:00:00')              │                    │
│    └──────────────────────────────────────┘                    │
│                                                                  │
│    HATUA 9.2: Zalisha ID mpya                                   │
│    ┌──────────────────────────────────────┐                    │
│    │ ID = 15 (auto-generated)             │                    │
│    └──────────────────────────────────────┘                    │
│                                                                  │
│    HATUA 9.3: Rudisha ID                                        │
│    ┌──────────────────────────────────────┐                    │
│    │ RETURNING id → 15                    │                    │
│    └──────────────────────────────────────┘                    │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ↓
┌─────────────────────────────────────────────────────────────────┐
│ 10. KURUDI NYUMA (Response Journey)                            │
│                                                                  │
│    Repository → Service → Controller                            │
│                                                                  │
│    roleID = 15                                                  │
│                                                                  │
│    Controller: wrappers.MessageResponse(                        │
│      c,                                                         │
│      http.StatusCreated,  // 201                               │
│      "Admin created successfully"                               │
│    )                                                            │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ↓
┌─────────────────────────────────────────────────────────────────┐
│ 11. HTTP RESPONSE                                               │
│                                                                  │
│    HTTP/1.1 201 Created                                         │
│    Content-Type: application/json                               │
│                                                                  │
│    {                                                            │
│      "status": 201,                                             │
│      "message": "Admin created successfully"                    │
│    }                                                            │
└─────────────────────────────────────────────────────────────────┘
```

---

## MUHTASARI WA MTIRIRIKO

### **Kwa Kifupi:**

1. **Entity** → Inaelezea muundo wa data na sheria za biashara
2. **UseCase/Service** → Inafanya mantiki ya biashara na kuunganisha tabaka
3. **Repository** → Inatengeneza SQL na kuwasiliana na database
4. **Model** → Inaelezea muundo wa JSON kwa API
5. **Controller** → Inapokea HTTP request na kusafisha data
6. **Route** → Inaelekeza request kwenye controller sahihi
7. **Database** → Inahifadhi data kabisa

### **Faida za Muundo Huu:**

✅ **Separation of Concerns** - Kila tabaka lina jukumu lake  
✅ **Testability** - Unaweza kujaribu kila tabaka peke yake  
✅ **Maintainability** - Rahisi kudumisha na kubadilisha  
✅ **Reusability** - Unaweza kutumia tena components  
✅ **Security** - Usalama katika tabaka nyingi  
✅ **Scalability** - Rahisi kupanua mfumo  

---

## HATUA ZA KUANZISHA PROGRAMU

### **HATUA 1: KUANZISHA PROGRAMU (`app.go`)**

```go
// Mstari 14-22: Kuanzisha logger
func init() {
    path := config.LoggerPath()              // Kupata njia ya kuhifadhi logs
    log.Infoln(path)                         // Kuonyesha njia
    log.SetOptions(                          // Kuweka chaguo za logger
        log.Development(),                    // Mode ya maendeleo
        log.WithCaller(true),                // Onyesha mahali log ilitoka
        log.WithLogDirs(path),               // Njia ya kuhifadhi logs
    )
}
```

**Maelezo:**
- Kabla programu haijaanza, `init()` inaitwa moja kwa moja
- Inaweka logger ili kurekodi matukio yote ya mfumo
- Logger itahifadhi kumbukumbu katika folda `.storage/.logs`

```go
// Mstari 23-46: Funguo kuu ya programu
func main() {
    database.Connect()                        // Mstari 25: Unganisha na database
    defer database.Close()                    // Mstari 26: Funga muunganisho baada ya programu kumalizika
    go webserver.StartWebserver()            // Mstari 27: Anzisha webserver katika goroutine nyingine
    
    defer os.Exit(0)                         // Mstari 29: Hakikisha programu inafunga vizuri
    
    stop := make(chan os.Signal, 1)          // Mstari 31: Tengeneza channel ya kupokea ishara
    
    signal.Notify(                           // Mstari 35-40: Sikiliza ishara za kufunga
        stop,
        syscall.SIGHUP,                      // Ishara ya kufunga terminal
        syscall.SIGINT,                      // Ctrl+C
        syscall.SIGQUIT,                     // Ctrl+\
    )
    <-stop                                   // Mstari 41: Subiri ishara ya kufunga
    log.Infoln("auth webserver is shutting down .... 👋 !")
}
```

**Maelezo:**
- **Mstari 25**: Inaunganisha na PostgreSQL database
- **Mstari 26**: `defer` inahakikisha muunganisho unafungwa hata kama kuna hitilafu
- **Mstari 27**: `go` inaanzisha webserver katika thread tofauti (concurrent)
- **Mstari 31-41**: Programu inasubiri ishara ya kufunga (kama Ctrl+C)

---

### **HATUA 2: KUUNGANISHA NA DATABASE (`services/database/database.go`)**

```go
// Mstari 12-14: Variables za global
var once sync.Once                           // Kuhakikisha database inaunganishwa mara moja tu
var instance *pgxpool.Pool                   // Instance ya muunganisho wa database
var err error                                // Hitilafu ikitokea

// Mstari 16-29: Funguo ya kuunganisha
func Connect() (*pgxpool.Pool, error) {
    once.Do(func() {                         // Mstari 17: Fanya kazi hii mara moja tu
        connectionString := config.GetDatabaseConnection()  // Mstari 18: Pata string ya muunganisho
        instance, err = pgxpool.Connect(context.Background(), connectionString)  // Mstari 19: Unganisha
        if err != nil {
            log.Errorf("unable to create a database instance")
        }
    })
    if err != nil {
        log.Errorf("unable to connect to database: %v\n", err)
        return nil, err
    }
    return instance, err                     // Mstari 28: Rudisha instance
}
```

**Maelezo:**
- **Mstari 17**: `sync.Once` inahakikisha muunganisho unatengenezwa mara moja tu (Singleton pattern)
- **Mstari 18**: Inapata taarifa za muunganisho kutoka `config.yml`:
  - Host: `127.0.0.1`
  - Database: `library`
  - User: `root`
  - Password: `Emmanuel@100`
  - Port: `5432`
- **Mstari 19**: Inaunganisha na PostgreSQL kwa kutumia pgx driver

---

### **HATUA 3: KUANZISHA WEBSERVER (`webserver/webserver.go`)**

```go
// Mstari 15-38: Kuanzisha webserver
func StartWebserver() {
    e := echo.New()                          // Mstari 17: Tengeneza instance ya Echo framework
    
    e.HideBanner = true                      // Mstari 20: Ficha banner ya Echo
    
    routes.Routers(e)                        // Mstari 23: Sajili njia zote za API
    
    helpers.Init()                           // Mstari 26: Anzisha cache
    systems.Init()                           // Mstari 27: Anzisha mifumo mingine
    
    e.Debug = true                           // Mstari 30: Washa mode ya debug
    
    cfg, err := config.New()                 // Mstari 32: Pata usanidi
    if err != nil {
        log.Errorf("error getting config: %v", err)
    }
    address := fmt.Sprintf("%v:%v", cfg.WebServer.PublicHost, cfg.WebServer.Port)  // Mstari 36
    e.Logger.Fatal(e.Start(address))        // Mstari 37: Anzisha server kwenye 127.0.0.1:4601
}
```

**Maelezo:**
- **Mstari 17**: Echo ni web framework rahisi na ya haraka kwa Go
- **Mstari 23**: Inasajili njia zote za API (routes)
- **Mstari 36-37**: Server inaanza kusikiliza kwenye `127.0.0.1:4601`

---

### **HATUA 4: KUSAJILI NJIA ZA API (`webserver/routes/router.go` & `api.go`)**

#### **router.go - Middleware za Jumla**

```go
// Mstari 14-34: Kusajili middleware na routes
func Routers(app *echo.Echo) {
    
    app.Use(middlewares.Cors())              // Mstari 17: Ruhusu CORS (Cross-Origin requests)
    app.Use(middlewares.Gzip())              // Mstari 18: Compress majibu
    app.Use(middlewares.Logger(true))        // Mstari 19: Rekodi kila request
    app.Use(middlewares.Secure())            // Mstari 20: Ongeza usalama headers
    app.Use(middlewares.Recover())           // Mstari 21: Zuia programu kuzimika kwa hitilafu
    
    app.Validator = validator.GetValidator() // Mstari 29: Weka validator ya custom
    
    ApiRouters(app)                          // Mstari 32: Sajili njia za API
    go generateRoutes(app)                   // Mstari 33: Tengeneza faili ya njia zote
}
```

**Maelezo:**
- **Middleware** ni kazi zinazofanyika kabla request haijafika kwa controller:
  - **CORS**: Inaruhusu maombi kutoka domain nyingine
  - **Gzip**: Inapunguza ukubwa wa data inayotumwa
  - **Logger**: Inarekodi kila request (wakati, njia, mtumiaji, n.k.)
  - **Secure**: Inaongeza HTTP headers za usalama
  - **Recover**: Inazuia server kuzimika kama kuna panic

#### **api.go - Njia za API**

```go
// Mstari 9-38: Kusajili endpoints za API
func ApiRouters(app *echo.Echo) {
    
    api := app.Group("library/api/v1")       // Mstari 16: Tengeneza kundi la njia
    
    // Njia za Role
    role := api.Group("/role")               // Mstari 23: Kundi la role
    {
        role.POST("/list", controllers.ListRole)         // Orodha ya roles
        role.POST("/show", controllers.GetRole)          // Pata role moja
        role.POST("/create", controllers.CreateRole)     // Tengeneza role mpya
        role.POST("/update", controllers.UpdateRole)     // Sasisha role
        role.POST("/delete", controllers.SoftDeleteRole) // Futa kwa soft delete
        role.POST("/destroy", controllers.DestroyRole)   // Futa kabisa
    }
    
    // Njia za Staff
    staff := api.Group("/staff")             // Mstari 33
    {
        staff.POST("/create", controllers.CreateStaff)   // Tengeneza mfanyakazi
    }
}
```

**Maelezo:**
- Njia zote zinaanza na prefix: `library/api/v1`
- Kwa mfano, kutengeneza role: `POST http://127.0.0.1:4601/library/api/v1/role/create`

---

### **HATUA 5: CONTROLLER INAPOKEA REQUEST (`webserver/controllers/role_controller.go`)**

Hebu tuangalie mfano wa **CreateRole** kwa undani:

```go
// Mstari 106-127: Controller ya kutengeneza role
func CreateRole(c echo.Context) error {
    // HATUA 5.1: Bind JSON data kutoka request
    m := &models.Role{}                      // Mstari 107: Tengeneza struct tupu
    if err := c.Bind(m); util.IsError(err) { // Mstari 108: Jaza data kutoka request body
        log.Errorf("error binding Role : %v", err)
        return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
    }
    
    // HATUA 5.2: Weka mtumiaji aliyetengeneza
    m.CreatedBy = 1                          // Mstari 113: ID ya mtumiaji (kwa sasa ni hardcoded)
    
    // HATUA 5.3: Thibitisha data
    customValidator, _ := c.Echo().Validator.(*validator.CustomValidator)  // Mstari 114
    if err := customValidator.ValidateStructPartial(m, "Name", "Description", "CreatedBy"); err != nil {
        log.Errorf("error validating Role  model: %v", err)
        return wrappers.ValidationErrorResponse(c, err)  // Rudisha hitilafu ya validation
    }
    
    // HATUA 5.4: Ita service layer
    service := role.NewService()             // Mstari 120: Tengeneza instance ya service
    _, err := service.Create(m.Name, m.Description, m.CreatedBy)  // Mstari 121: Ita funguo ya kutengeneza
    if util.IsError(err) {
        log.Errorf("error creating new %v: %v", m.Name, err)
        return wrappers.ErrorResponse(c, http.StatusInternalServerError, internalServerErrorMsg)
    }
    
    // HATUA 5.5: Rudisha jibu la mafanikio
    return wrappers.MessageResponse(c, http.StatusCreated, m.Name+" created successfully")  // Mstari 126
}
```

**Maelezo kwa Kina:**

**Mstari 107-111 - Binding Data:**
- `models.Role{}` ni struct inayoelezea muundo wa data:
  ```go
  type Role struct {
      ID          int32     `json:"id"`
      Name        string    `json:"name" validate:"required"`
      Description string    `json:"description" validate:"required"`
      CreatedBy   int32     `json:"created_by" validate:"numeric,required"`
  }
  ```
- `c.Bind(m)` inabadilisha JSON kutoka request kuwa Go struct
- Kwa mfano, JSON hii:
  ```json
  {
    "name": "Admin",
    "description": "Administrator role"
  }
  ```
  Inabadilishwa kuwa struct ya Go

**Mstari 113 - CreatedBy:**
- Inaweka ID ya mtumiaji aliyetengeneza rekodi
- Kwa sasa ni hardcoded (1), lakini kawaida inapata kutoka JWT token

**Mstari 114-118 - Validation:**
- Inathibitisha kwamba:
  - `Name` haiko tupu
  - `Description` haiko tupu
  - `CreatedBy` ni nambari na haiko tupu
- Kama validation imeshindwa, inarudisha hitilafu 400 (Bad Request)

**Mstari 120-125 - Service Layer:**
- Inatengeneza instance ya `role.Service`
- Inaita `service.Create()` kupitisha data kwenye tabaka la biashara

**Mstari 126 - Response:**
- Inarudisha jibu la mafanikio na status code 201 (Created)

---

### **HATUA 6: SERVICE/USECASE LAYER (`services/usecase/role/role_service.go`)**

```go
// Mstari 20-31: Mantiki ya biashara ya kutengeneza role
func (s *Service) Create(name string, description string, createdBy int32) (int32, error) {
    // HATUA 6.1: Tengeneza entity na thibitisha
    role, err := entity.NewRole(name, description, createdBy)  // Mstari 21
    if err != nil {
        return role.ID, err              // Mstari 23: Rudisha hitilafu ya validation
    }
    
    // HATUA 6.2: Hifadhi kwenye database kupitia repository
    roleID, err := s.repo.Create(role)   // Mstari 25
    if err != nil {
        return role.ID, err
    }
    
    // HATUA 6.3: Rudisha ID ya role iliyotengenezwa
    return roleID, err                   // Mstari 29
}
```

**Maelezo:**
- **Mstari 21**: Inatengeneza entity ya `Role` na kuithibitisha
- **Mstari 25**: Inapitisha entity kwenye repository layer kwa kuhifadhi
- Tabaka hili linaweza kuongeza mantiki ya biashara kama:
  - Kuthibitisha kwamba jina la role halijasajiliwa
  - Kutuma email baada ya kutengeneza
  - Kurekodi audit trail

---

### **HATUA 7: ENTITY LAYER (`services/entity/role_entity.go`)**

```go
// Mstari 9-19: Muundo wa Entity
type Role struct {
    ID          int32
    Name        string
    Description string
    CreatedBy   int32
    CreatedAt   time.Time
    UpdatedBy   int32
    UpdatedAt   time.Time
    DeletedBy   int32
    DeletedAt   time.Time
}

// Mstari 21-33: Constructor ya kutengeneza role mpya
func NewRole(name, description string, createdBy int32) (*Role, error) {
    role := &Role{                       // Mstari 22: Tengeneza instance
        Name:        name,
        Description: description,
        CreatedBy:   createdBy,
    }
    err := role.ValidateCreate()         // Mstari 27: Thibitisha data
    if err != nil {
        log.Errorf("error validating Role entity %v", err)
        return nil, err
    }
    return role, nil                     // Mstari 32: Rudisha entity
}

// Mstari 35-46: Uthibitishaji wa data
func (r *Role) ValidateCreate() error {
    if r.Name == "" {                    // Mstari 36: Thibitisha jina
        return errors.New("error validating Role entity, name field required")
    }
    if r.Description == "" {             // Mstari 39: Thibitisha maelezo
        return errors.New("error validating Role entity, description field required")
    }
    if r.CreatedBy <= 0 {                // Mstari 42: Thibitisha mtumiaji
        return errors.New("error validating Role entity, created_by field required")
    }
    return nil                           // Mstari 45: Data ni sahihi
}
```

**Maelezo:**
- **Entity** ni muundo wa data wa ndani wa mfumo
- Inatofautiana na `models.Role` (inayotumika kwa API)
- Ina mantiki ya uthibitishaji wa biashara
- Inaweza kuwa na methods za ziada za kihesabu

---

### **HATUA 8: REPOSITORY LAYER (`services/repository/role_repository.go`)**

```go
// Mstari 20-32: Muundo wa Repository
type RoleConn struct {
    conn *pgxpool.Pool                   // Muunganisho wa database
}

func NewRole() *RoleConn {
    conn, err := database.Connect()      // Mstari 25: Pata muunganisho wa database
    if util.IsError(err) {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
    }
    return &RoleConn{                    // Mstari 29: Rudisha instance
        conn: conn,
    }
}

// Mstari 50-61: Funguo ya kuhifadhi kwenye database
func (con *RoleConn) Create(e *entity.Role) (int32, error) {
    var RoleID int32                     // Mstari 51: Variable ya ID itakayorudishwa
    
    // Mstari 52: SQL query ya INSERT
    query := `INSERT INTO roles (name, description, created_by, created_at) 
              VALUES ($1, $2, $3, $4) 
              RETURNING id`
    
    // Mstari 53: Tekeleza query
    err := con.conn.QueryRow(
        context.Background(),            // Context ya Go
        query,                           // SQL query
        e.Name,                          // $1 - Jina la role
        e.Description,                   // $2 - Maelezo
        e.CreatedBy,                     // $3 - ID ya mtumiaji
        time.Now(),                      // $4 - Wakati wa sasa
    ).Scan(&RoleID)                      // Pokea ID iliyorudishwa
    
    if util.IsError(err) {
        if err.Error() == error_message.ErrNoResultSet.Error() {
            return RoleID, nil
        }
        log.Errorf("error creating role from table %v: %v", roleTableName, err)
    }
    return RoleID, err                   // Mstari 60: Rudisha ID
}
```

**Maelezo kwa Kina:**

**Mstari 52 - SQL Query:**
- `INSERT INTO roles` - Ingiza data kwenye jedwali la roles
- `VALUES ($1, $2, $3, $4)` - Placeholders za usalama (prepared statement)
- `RETURNING id` - PostgreSQL inarudisha ID ya rekodi iliyotengenezwa

**Mstari 53 - QueryRow:**
- `context.Background()` - Context ya Go kwa ajili ya timeout na cancellation
- Parameters (`e.Name`, `e.Description`, etc.) zinabadilisha `$1, $2, $3, $4`
- `.Scan(&RoleID)` - Pokea ID kutoka `RETURNING id`

**Usalama:**
- Prepared statements zinazuia SQL injection
- Kila parameter inasafishwa na driver ya database

---

## MTIRIRIKO KAMILI WA DATA - MFANO WA CREATE ROLE

### **Request kutoka Mtumiaji:**
```http
POST http://127.0.0.1:4601/library/api/v1/role/create
Content-Type: application/json

{
  "name": "Admin",
  "description": "Administrator role with full access"
}
```

### **Mtiririko wa Data:**

```
1. REQUEST INAPOKEA
   ↓
   [Middleware: CORS] → Thibitisha origin ya request
   ↓
   [Middleware: Logger] → Rekodi: "POST /library/api/v1/role/create"
   ↓
   [Middleware: Secure] → Ongeza security headers
   ↓
   [Middleware: Recover] → Weka panic handler

2. ROUTING
   ↓
   [routes/api.go:27] → Pata controller: controllers.CreateRole
   ↓

3. CONTROLLER LAYER
   ↓
   [role_controller.go:108] → Bind JSON → models.Role struct
   ↓
   [role_controller.go:113] → Weka CreatedBy = 1
   ↓
   [role_controller.go:115] → Validate: Name ✓, Description ✓, CreatedBy ✓
   ↓
   [role_controller.go:120] → Tengeneza service = role.NewService()
   ↓
   [role_controller.go:121] → Ita service.Create("Admin", "Administrator...", 1)

4. SERVICE/USECASE LAYER
   ↓
   [role_service.go:21] → entity.NewRole("Admin", "Administrator...", 1)
   ↓

5. ENTITY LAYER
   ↓
   [role_entity.go:22-26] → Tengeneza Role struct
   ↓
   [role_entity.go:27] → ValidateCreate()
   ↓
   [role_entity.go:36-45] → Thibitisha: Name ✓, Description ✓, CreatedBy ✓
   ↓
   [role_entity.go:32] → Rudisha entity ya Role
   ↓

6. KURUDI KWA SERVICE LAYER
   ↓
   [role_service.go:25] → s.repo.Create(role)
   ↓

7. REPOSITORY LAYER
   ↓
   [role_repository.go:25] → Pata database connection
   ↓
   [role_repository.go:52-53] → Tengeneza SQL query:
                                 INSERT INTO roles (name, description, created_by, created_at)
                                 VALUES ('Admin', 'Administrator...', 1, '2024-02-15 22:00:00')
                                 RETURNING id
   ↓

8. DATABASE (PostgreSQL)
   ↓
   [PostgreSQL] → Tekeleza INSERT statement
   ↓
   [PostgreSQL] → Tengeneza ID mpya (kwa mfano: 15)
   ↓
   [PostgreSQL] → RETURN id = 15
   ↓

9. KURUDI NYUMA
   ↓
   [role_repository.go:60] → Rudisha roleID = 15
   ↓
   [role_service.go:29] → Rudisha roleID = 15
   ↓
   [role_controller.go:126] → Tengeneza response:
                               {
                                 "status": 201,
                                 "message": "Admin created successfully"
                               }
   ↓

10. RESPONSE KWA MTUMIAJI
   ↓
   [Middleware: Gzip] → Compress response
   ↓
   [Middleware: Logger] → Rekodi: "201 Created - 45ms"
   ↓
   HTTP/1.1 201 Created
   Content-Type: application/json
   
   {
     "status": 201,
     "message": "Admin created successfully"
   }
```

---

## MIFANO YA OPERATIONS NYINGINE

### **1. LIST ROLES (Orodha ya Roles)**

**Request:**
```http
POST http://127.0.0.1:4601/library/api/v1/role/list
Content-Type: application/json

{
  "page": 1,
  "page_size": 10,
  "sort_by": "name",
  "sort_order": "ASC"
}
```

**Mtiririko:**
```
Controller [ListRole] 
  ↓
Service [List(filter)]
  ↓
Repository [List(filter)]
  ↓
SQL: SELECT id, name, description, created_by, created_at, updated_by, updated_at
     FROM roles 
     WHERE deleted_at IS NULL 
     ORDER BY name ASC 
     LIMIT 10 OFFSET 0
  ↓
Database → Rudisha orodha ya roles
  ↓
Response: {
  "status": 200,
  "data": [
    {"id": 1, "name": "Admin", "description": "..."},
    {"id": 2, "name": "User", "description": "..."}
  ],
  "meta": {
    "page": 1,
    "page_size": 10,
    "total_count": 25,
    "total_pages": 3
  }
}
```

---

### **2. UPDATE ROLE (Sasisha Role)**

**Request:**
```http
POST http://127.0.0.1:4601/library/api/v1/role/update
Content-Type: application/json

{
  "id": 15,
  "name": "Super Admin",
  "description": "Super administrator with all privileges"
}
```

**Mtiririko:**
```
Controller [UpdateRole]
  ↓
  Validate: ID ✓, Name ✓, Description ✓
  ↓
Service [Update(entity)]
  ↓
Repository [Update(entity)]
  ↓
SQL: UPDATE roles 
     SET name = 'Super Admin', 
         description = 'Super administrator...', 
         updated_by = 1, 
         updated_at = '2024-02-15 22:05:00' 
     WHERE id = 15
  ↓
Database → Tekeleza UPDATE
  ↓
Response: {
  "status": 202,
  "message": "Super Admin updated successfully"
}
```

---

### **3. SOFT DELETE ROLE (Futa kwa Unyenyekevu)**

**Request:**
```http
POST http://127.0.0.1:4601/library/api/v1/role/delete
Content-Type: application/json

{
  "id": 15
}
```

**Mtiririko:**
```
Controller [SoftDeleteRole]
  ↓
Service [SoftDelete(id, deletedBy)]
  ↓
Repository [SoftDelete(id, deletedBy)]
  ↓
SQL: UPDATE roles 
     SET deleted_by = 1, 
         deleted_at = '2024-02-15 22:10:00' 
     WHERE id = 15
  ↓
Database → Weka deleted_at (rekodi haijafutwa kabisa)
  ↓
Response: {
  "status": 202,
  "message": "record deleted successfully"
}
```

**Kumbuka:** Soft delete inahifadhi data lakini inaificha kutoka kwenye queries za kawaida.

---

### **4. HARD DELETE ROLE (Futa Kabisa)**

**Request:**
```http
POST http://127.0.0.1:4601/library/api/v1/role/destroy
Content-Type: application/json

{
  "id": 15
}
```

**Mtiririko:**
```
Controller [DestroyRole]
  ↓
Service [HardDelete(id)]
  ↓
Repository [HardDelete(id)]
  ↓
SQL: DELETE FROM roles WHERE id = 15
  ↓
Database → Futa rekodi kabisa (haiwezi kurejeshwa)
  ↓
Response: {
  "status": 202,
  "message": "record deleted successfully"
}
```

---

## MUUNDO WA ARCHITECTURE (Clean Architecture)

```
┌─────────────────────────────────────────────────────────┐
│                    PRESENTATION LAYER                    │
│  (Controllers, Routes, Middlewares, Request/Response)   │
│                                                          │
│  - Kupokea HTTP requests                                │
│  - Validation ya input                                  │
│  - Kutuma responses                                     │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│                   BUSINESS LOGIC LAYER                   │
│              (UseCase/Service, Entity)                  │
│                                                          │
│  - Mantiki ya biashara                                  │
│  - Validation ya business rules                         │
│  - Orchestration ya operations                          │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│                     DATA ACCESS LAYER                    │
│                    (Repository)                         │
│                                                          │
│  - Mawasiliano na database                              │
│  - SQL queries                                          │
│  - Data mapping                                         │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│                       DATABASE                           │
│                   (PostgreSQL)                          │
│                                                          │
│  - Kuhifadhi data                                       │
│  - Transactions                                         │
│  - Data integrity                                       │
└─────────────────────────────────────────────────────────┘
```

---

## FAIDA ZA MUUNDO HUU

### **1. Separation of Concerns (Kugawanya Majukumu)**
- Kila tabaka lina jukumu lake maalum
- Controller haijui kuhusu database
- Repository haijui kuhusu HTTP

### **2. Testability (Uwezo wa Kujaribu)**
- Unaweza kujaribu kila tabaka peke yake
- Mock dependencies kwa testing

### **3. Maintainability (Urahisi wa Kudumisha)**
- Mabadiliko kwenye database hayaathiri controller
- Unaweza kubadilisha database bila kubadilisha business logic

### **4. Reusability (Kutumia Tena)**
- Service layer inaweza kutumika na API nyingine (GraphQL, gRPC)
- Repository inaweza kutumika na services nyingine

### **5. Security (Usalama)**
- Prepared statements zinazuia SQL injection
- Validation katika tabaka nyingi
- Middleware za usalama

---

## TEKNOLOJIA ZINAZOTUMIKA

### **1. Echo Framework**
- Web framework ya haraka na rahisi
- Routing, middleware, request binding
- Website: https://echo.labstack.com/

### **2. pgx/pgxpool**
- PostgreSQL driver ya Go
- Connection pooling
- Performance ya juu

### **3. Viper**
- Configuration management
- Kusoma config.yml

### **4. Zap Logger**
- High-performance logging
- Structured logging

### **5. Go Validator**
- Validation ya struct fields
- Custom validation rules

---

## JINSI YA KUANZISHA MRADI

### **1. Mahitaji:**
```bash
- Go 1.25 au zaidi
- PostgreSQL 12 au zaidi
```

### **2. Install Dependencies:**
```bash
go mod download
```

### **3. Sanidi Database:**
```sql
CREATE DATABASE library;
```

### **4. Badilisha config.yml:**
```yaml
database:
  name: library
  user: your_username
  password: your_password
  port: 5432
```

### **5. Run Migrations:**
```bash
# Tekeleza SQL scripts katika folda migrations/
```

### **6. Anzisha Server:**
```bash
go run app.go
```

### **7. Jaribu API:**
```bash
curl -X POST http://127.0.0.1:4601/library/api/v1/role/create \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Admin",
    "description": "Administrator role"
  }'
```

---

## MUHTASARI WA MTIRIRIKO

**Kwa ufupi, hii ndiyo safari ya data kutoka request hadi database:**

1. **Mtumiaji** anatuma HTTP request
2. **Middleware** zinasafisha na kurekodi request
3. **Router** inapeleka request kwa controller sahihi
4. **Controller** inapokea data, inaithibitisha, na kuita service
5. **Service** inafanya mantiki ya biashara na kuita repository
6. **Repository** inatengeneza SQL query na kuwasiliana na database
7. **Database** inahifadhi data na kurudisha matokeo
8. **Response** inarudi nyuma kupitia tabaka zote hadi kwa mtumiaji

Kila tabaka lina jukumu lake maalum, na hii inafanya mfumo kuwa rahisi kudumisha na kupanua!

---

## MAWASILIANO

Kwa maswali au usaidizi, wasiliana na timu ya maendeleo.

**Asante kwa kutumia mfumo huu!** 🚀
