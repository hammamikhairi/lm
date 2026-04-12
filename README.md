# lm - Refactoring Guide (GOF + SOLID + GRASP + OCL)

This project is a CLI directory lister (similar to `ls`) written in Go.

The current implementation works, but most logic is in `main` and global functions:

- `main.go`: argument parsing, terminal setup, file traversal, output orchestration
- `format.go`: size/date/mode/name formatting + icon selection calls
- `icons.go`: icon/color registry

This README provides a clean, modular plan for a 5-student team where each student implements:

- 1 GOF Creational pattern
- 1 GOF Structural pattern
- 1 GOF Behavioral pattern
- 1 SOLID principle
- 1 GRASP principle
- 1 OCL constraint

## 1) Project Analysis: Where Patterns Fit

### Current pain points

1. Too much responsibility in `main.go` (violates SRP).
2. Output rendering and business logic are tightly coupled.
3. Icon detection is currently map/if-based and not extensible.
4. Sorting and filtering behavior is hard-coded.
5. Global state (`dirPath`, `termGap`) makes testing difficult.

### Natural extension points

1. File scanning and metadata acquisition.
2. Sorting behavior (name, size, date, type).
3. Output rendering (compact table, long format, JSON).
4. Icon resolution chain (special names, tests, extension, fallback).
5. Terminal-dependent behavior (width, color, spacing).

## 2) Target Modular Architecture

Use this package structure during refactoring:

```text
lm/
  cmd/
	 lm/
		main.go
  internal/
	 app/
		controller.go
		config.go
	 domain/
		entry.go
		constraints.go
	 scanner/
		scanner.go
		os_scanner.go
	 sort/
		strategy.go
		by_name.go
		by_size.go
		by_date.go
	 render/
		renderer.go
		long_renderer.go
		json_renderer.go
		decorator_color.go
		decorator_icon.go
	 icon/
		resolver.go
		chain.go
		registry.go
	 factory/
		formatter_factory.go
	 observer/
		events.go
		logger_observer.go
```

## 3) Team Distribution (5 Students)

Each student gets one implementation set. Patterns are chosen to fit your current code.

| Student | Creational              | Structural | Behavioral              | SOLID Focus | GRASP Focus                | OCL Focus                    |
| ------- | ----------------------- | ---------- | ----------------------- | ----------- | -------------------------- | ---------------------------- |
| S1      | Factory Method          | Adapter    | Strategy                | SRP         | Controller                 | Entry size/date invariants   |
| S2      | Abstract Factory        | Decorator  | Template Method         | OCP         | Information Expert         | Renderer postconditions      |
| S3      | Builder                 | Facade     | Command                 | DIP         | Pure Fabrication           | Command preconditions        |
| S4      | Singleton (config only) | Proxy      | Observer                | ISP         | Indirection                | Event consistency invariants |
| S5      | Prototype               | Composite  | Chain of Responsibility | LSP         | Low Coupling/High Cohesion | Icon resolution constraints  |

## 4) GOF Patterns: What to Implement Here

### A) Creational

1. Factory Method
   - Create `EntryFactory` that builds `FileEntry` from `os.DirEntry` + `os.FileInfo`.
   - Good location: `internal/factory` + `internal/domain`.

2. Abstract Factory
   - Build `RendererFactory` for different output families (plain, color, json).
   - Each family returns compatible formatter components.

3. Builder
   - Build a `LineBuilder` for formatted output line parts:
     mode, date, size, icon, filename.

4. Singleton
   - Keep this limited to immutable runtime config (`AppConfig`).
   - Avoid singleton business services.

5. Prototype
   - Clone base render styles or icon themes for quick variants.

### B) Structural

1. Adapter
   - Wrap OS-specific terminal width logic (`unix.IoctlGetWinsize`) behind `TerminalInfoProvider`.

2. Decorator
   - Decorate base renderer with optional features:
     icon decorator, color decorator, executable marker decorator.

3. Facade
   - `ListFacade` exposes one simple method:
     `List(path, options) -> []RenderedLine`.

4. Proxy
   - Cache expensive metadata reads through `FileInfoProxy`.

5. Composite
   - Represent tree listing with `DirectoryNode` containing child `EntryNode`s.

### C) Behavioral

1. Strategy
   - `SortStrategy` (`ByName`, `BySize`, `ByDate`).
   - `IconResolveStrategy` if needed.

2. Template Method
   - Define invariant listing flow in abstract base workflow:
     scan -> validate -> sort -> render.
   - Override sort/render steps in concrete listers.

3. Command
   - CLI actions as commands: `ListCommand`, `ListTreeCommand`, `ListJsonCommand`.

4. Observer
   - Emit events (`ScanStarted`, `ScanFinished`, `RenderDone`) for logging/stats.

5. Chain of Responsibility
   - Icon resolution chain:
     special-name handler -> test-file handler -> extension handler -> unknown handler.

## 5) SOLID Principles: Concrete Application

1. SRP
   - Separate responsibilities:
     scanner, sorter, renderer, icon resolver, controller.

2. OCP
   - New sort mode or renderer should be added by implementing interfaces,
     not by editing existing switch/if blocks.

3. LSP
   - Any `Renderer` implementation must work where `Renderer` is expected.
   - Ensure substitutability for decorators/proxies.

4. ISP
   - Split interfaces:
     `Scanner`, `Sorter`, `Renderer`, `IconResolver` instead of one large interface.

5. DIP
   - `Controller` depends on interfaces, not concrete `os` implementations.

## 6) GRASP Principles: Concrete Application

1. Controller
   - `ListingController` handles request orchestration from `main`.

2. Information Expert
   - `FileEntry` calculates behaviors related to its own metadata.

3. Low Coupling
   - Use dependency injection in `main` to wire modules.

4. High Cohesion
   - Keep each package focused (sorting code only in `sort/`, etc.).

5. Indirection
   - Interfaces between controller and infra services (`Scanner`, `TerminalInfoProvider`).

6. Pure Fabrication
   - `IconRegistryService` as a dedicated service, not attached to unrelated entities.

## 7) OCL Constraints (Model-Level Contracts)

Use OCL as design-by-contract documentation. Even if Go does not execute OCL directly,
keep these constraints in `docs/constraints.ocl` and enforce with tests.

### Suggested model elements

- `FileEntry(name: String, size: Integer, isDir: Boolean, mode: String, modifiedAt: DateTime)`
- `RenderLine(text: String, visibleLength: Integer)`
- `IconResult(icon: String, color: String, source: String)`

### Example OCL constraints

```ocl
context FileEntry
inv NonNegativeSize: self.size >= 0

context FileEntry
inv DirectorySizePolicy: self.isDir implies self.size = 0

context FileEntry
inv ValidName: self.name.size() > 0

context IconResult
inv NonEmptyIcon: self.icon.size() > 0

context IconResult
inv ValidSource: Set{'SPECIAL','TEST','EXT','DEFAULT'}->includes(self.source)

context RenderLine
inv VisibleLengthNonNegative: self.visibleLength >= 0

context ListingController::list(path: String)
pre PathProvided: path.size() > 0

context ListingController::list(path: String)
post ResultNotNull: result <> null
```

## 8) Suggested Implementation Roadmap

1. Extract interfaces and move orchestration from `main.go` to `ListingController`.
2. Implement `Strategy` for sorting.
3. Implement `Chain of Responsibility` for icon resolution.
4. Add `Decorator` around rendering pipeline.
5. Add one creational pattern (`Factory Method` or `Abstract Factory`) for assembly.
6. Add OCL file and corresponding Go unit tests for invariants/pre/postconditions.

## 9) Deliverables Checklist (Per Student)

Each student should submit:

1. Pattern implementation code.
2. Short UML/class diagram for their patterns.
3. One paragraph explaining SOLID + GRASP choice.
4. One OCL constraint file section + test proving enforcement.
5. Before/after explanation of extensibility improvement.

## 10) Evaluation Tips

1. Prefer small interfaces and constructor injection.
2. Avoid adding global variables.
3. Keep old behavior identical first, then extend features.
4. Add tests before and after each major refactor.

---

If you want, the next step can be a concrete task breakdown with exact file-by-file assignments for S1..S5 and starter Go interfaces for each pattern.
