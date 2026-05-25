# Storage Format Overview

Confluence uses a specific XHTML dialect called "Storage Format" for page content. This server uses structured blocks as an intermediate representation (IR) to safely work with this format.

## Why Structured Blocks?

### The Problem

When AI assistants work directly with Confluence XHTML:

1. **Wrong format** - LLMs generate HTML5 or Markdown, not Storage Format
2. **Table corruption** - Missing `<tbody>`, incorrect `<thead>` usage
3. **Broken macros** - Missing `ac:` namespaces
4. **Data loss** - Round-tripping through incorrect formats

### The Solution

Structured blocks solve these problems:

1. **JSON format** - LLMs produce JSON reliably
2. **Safe rendering** - Go code generates valid XHTML
3. **Validation** - Errors caught before API calls
4. **Reversible** - Parse XHTML back to blocks

## Architecture

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│   AI Assistant  │────▶│  Structured IR  │────▶│  Storage XHTML  │
│   (JSON input)  │     │   (Go types)    │     │   (validated)   │
└─────────────────┘     └─────────────────┘     └─────────────────┘
                               │
                               ▼
                        ┌─────────────────┐
                        │  Confluence API │
                        └─────────────────┘
```

## Block Types

| Type | Description |
|------|-------------|
| `Paragraph` | Text paragraph |
| `Heading` | H1-H6 headings |
| `Table` | Tables with headers and rows |
| `BulletList` | Unordered lists |
| `NumberedList` | Ordered lists |
| `Macro` | Confluence macros (status, info, code) |
| `CodeBlock` | Code blocks with language |
| `HorizontalRule` | Horizontal dividers |

See [Block Types](blocks.md) for detailed documentation.

## Example Workflow

### 1. AI receives task

"Add a status table to page 12345"

### 2. AI produces JSON

```json
{
  "blocks": [
    {"type": "heading", "level": 2, "text": "Service Status"},
    {
      "type": "table",
      "headers": ["Service", "Status"],
      "rows": [
        ["API", {"macro": {"name": "status", "params": {"colour": "Green", "title": "OK"}}}],
        ["Auth", {"macro": {"name": "status", "params": {"colour": "Green", "title": "OK"}}}]
      ]
    }
  ]
}
```

### 3. Server renders XHTML

```xml
<h2>Service Status</h2>
<table>
  <tbody>
    <tr>
      <th>Service</th>
      <th>Status</th>
    </tr>
    <tr>
      <td>API</td>
      <td><ac:structured-macro ac:name="status">
        <ac:parameter ac:name="colour">Green</ac:parameter>
        <ac:parameter ac:name="title">OK</ac:parameter>
      </ac:structured-macro></td>
    </tr>
    ...
  </tbody>
</table>
```

### 4. Server validates

- Checks for forbidden tags (`<thead>`, `<script>`)
- Verifies `<tbody>` structure
- Validates macro namespaces

### 5. Updates Confluence

Valid XHTML is sent to the Confluence API.
