# Block Types

This page documents all supported block types for structured content.

## Paragraph

Basic text content.

**JSON:**

```json
{"type": "paragraph", "text": "This is a paragraph."}
```

**Renders to:**

```xml
<p>This is a paragraph.</p>
```

## Heading

H1-H6 headings.

**JSON:**

```json
{"type": "heading", "level": 2, "text": "Section Title"}
```

**Renders to:**

```xml
<h2>Section Title</h2>
```

## Table

Tables with headers and rows.

**JSON:**

```json
{
  "type": "table",
  "headers": ["Name", "Role", "Status"],
  "rows": [
    ["Alice", "Lead", "Active"],
    ["Bob", "Developer", "Active"]
  ]
}
```

**Renders to:**

```xml
<table>
  <tbody>
    <tr>
      <th>Name</th>
      <th>Role</th>
      <th>Status</th>
    </tr>
    <tr>
      <td>Alice</td>
      <td>Lead</td>
      <td>Active</td>
    </tr>
    <tr>
      <td>Bob</td>
      <td>Developer</td>
      <td>Active</td>
    </tr>
  </tbody>
</table>
```

**Note:** Headers are rendered as `<th>` in the first row of `<tbody>`. Confluence Storage Format does NOT support `<thead>`.

### Tables with Macros

Cells can contain macros:

```json
{
  "type": "table",
  "headers": ["Service", "Status"],
  "rows": [
    ["API", {"macro": {"name": "status", "params": {"colour": "Green", "title": "OK"}}}]
  ]
}
```

## Bullet List

Unordered lists.

**JSON:**

```json
{"type": "bullet_list", "items": ["First item", "Second item", "Third item"]}
```

**Renders to:**

```xml
<ul>
  <li>First item</li>
  <li>Second item</li>
  <li>Third item</li>
</ul>
```

## Numbered List

Ordered lists.

**JSON:**

```json
{"type": "numbered_list", "items": ["Step one", "Step two", "Step three"]}
```

**Renders to:**

```xml
<ol>
  <li>Step one</li>
  <li>Step two</li>
  <li>Step three</li>
</ol>
```

## Macro

Confluence macros.

### Status Macro

**JSON:**

```json
{
  "type": "macro",
  "name": "status",
  "params": {
    "colour": "Green",
    "title": "OK"
  }
}
```

**Renders to:**

```xml
<ac:structured-macro ac:name="status">
  <ac:parameter ac:name="colour">Green</ac:parameter>
  <ac:parameter ac:name="title">OK</ac:parameter>
</ac:structured-macro>
```

**Available colors:** Grey, Red, Yellow, Green, Blue

### Info Macro

**JSON:**

```json
{
  "type": "macro",
  "name": "info",
  "params": {"title": "Note"},
  "body": "This is important information."
}
```

### Warning Macro

**JSON:**

```json
{
  "type": "macro",
  "name": "warning",
  "params": {"title": "Warning"},
  "body": "Be careful!"
}
```

### Code Macro

**JSON:**

```json
{
  "type": "macro",
  "name": "code",
  "params": {"language": "python"},
  "body": "print('Hello, World!')"
}
```

## Code Block

Shorthand for code macro.

**JSON:**

```json
{"type": "code_block", "language": "go", "code": "fmt.Println(\"Hello\")"}
```

**Renders to:**

```xml
<ac:structured-macro ac:name="code">
  <ac:parameter ac:name="language">go</ac:parameter>
  <ac:plain-text-body><![CDATA[fmt.Println("Hello")]]></ac:plain-text-body>
</ac:structured-macro>
```

## Horizontal Rule

Divider line.

**JSON:**

```json
{"type": "horizontal_rule"}
```

**Renders to:**

```xml
<hr/>
```
