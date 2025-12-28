# Markdown Syntax Guide

## markdown headers
**Tags**: markdown, headers, headings, title, h1, h2, h3
**Keywords**: headers headings title h1 h2 h3 h4 h5 h6

Create headers

```markdown
# H1 Header
## H2 Header
### H3 Header
#### H4 Header
##### H5 Header
###### H6 Header
```

## markdown emphasis
**Tags**: markdown, emphasis, bold, italic, formatting
**Keywords**: emphasis bold italic strong formatting text style

Text formatting

```markdown
**bold text**
*italic text*
***bold and italic***
~~strikethrough~~
`inline code`
```

## markdown lists
**Tags**: markdown, lists, ordered, unordered, bullets
**Keywords**: lists ordered unordered bullets numbered items

Create lists

```markdown
# Unordered list
- Item 1
- Item 2
  - Nested item
  - Another nested

# Ordered list
1. First item
2. Second item
   1. Nested numbered
   2. Another nested

# Task list
- [x] Completed task
- [ ] Incomplete task
```

## markdown links
**Tags**: markdown, links, hyperlinks, urls, references
**Keywords**: links hyperlinks urls references href anchor

Create links

```markdown
[Link text](https://example.com)
[Link with title](https://example.com "Link Title")

# Reference-style links
[Link text][reference]
[reference]: https://example.com

# Auto-links
<https://example.com>
<email@example.com>
```

## markdown images
**Tags**: markdown, images, pictures, media, graphics
**Keywords**: images pictures media graphics photos embed

Embed images

```markdown
![Alt text](image.jpg)
![Alt text](image.jpg "Image Title")

# Reference-style images
![Alt text][image-reference]
[image-reference]: image.jpg

# With links
[![Alt text](image.jpg)](https://example.com)
```

## markdown code blocks
**Tags**: markdown, code, syntax, highlighting, fenced, blocks
**Keywords**: code syntax highlighting fenced blocks backticks triple

Code blocks with syntax highlighting

````markdown
```javascript
function hello() {
  console.log("Hello, World!");
}
```

```python
def hello():
    print("Hello, World!")
```

```bash
echo "Hello, World!"
```
````

## markdown blockquotes
**Tags**: markdown, blockquotes, quotes, citation
**Keywords**: blockquotes quotes citation reference indented

Create blockquotes

```markdown
> This is a blockquote
> It can span multiple lines

> Multi-paragraph quote
>
> Second paragraph

> Nested quote
>> Nested deeper
```

## markdown tables
**Tags**: markdown, tables, grid, data, tabular
**Keywords**: tables grid data tabular rows columns alignment

Create tables

```markdown
| Header 1 | Header 2 | Header 3 |
|----------|----------|----------|
| Row 1    | Data     | More     |
| Row 2    | Data     | More     |

# Alignment
| Left | Center | Right |
|:-----|:------:|------:|
| L    |   C    |     R |
```

## markdown horizontal rule
**Tags**: markdown, horizontal, rule, separator, divider, line
**Keywords**: horizontal rule separator divider line hr break

Create horizontal rules

```markdown
---
***
___
```

## markdown footnotes
**Tags**: markdown, footnotes, references, citations, notes
**Keywords**: footnotes references citations notes superscript

Add footnotes

```markdown
Here's a sentence with a footnote[^1].

[^1]: This is the footnote content.

Multiple references[^note] to the same note[^note].

[^note]: Shared footnote.
```

## markdown definition lists
**Tags**: markdown, definition, lists, terms, glossary
**Keywords**: definition lists terms glossary dictionary

Create definition lists

```markdown
Term 1
: Definition for term 1

Term 2
: First definition for term 2
: Second definition for term 2
```

## markdown html
**Tags**: markdown, html, raw, inline, embedded
**Keywords**: html raw inline embedded tags elements

Embed raw HTML

```markdown
<div style="color: red;">
  This is red text
</div>

<details>
  <summary>Click to expand</summary>
  Hidden content here
</details>

<kbd>Ctrl</kbd> + <kbd>C</kbd>
```

## markdown escaping
**Tags**: markdown, escape, special, characters, backslash
**Keywords**: escape special characters backslash literal symbols

Escape special characters

```markdown
\* Not a bullet point
\# Not a header
\[Not a link\]
\`Not code\`

Use backslash \ before:
* \* \_ \{ \} \[ \] \( \)
* \# \+ \- \. \!
* \| \~
```

## markdown github extensions
**Tags**: markdown, github, gfm, extensions, flavored
**Keywords**: github gfm extensions flavored syntax special

GitHub Flavored Markdown features

```markdown
# Task lists
- [x] Complete
- [ ] Incomplete

# Username mentions
@username

# Issue references
#123
owner/repo#123

# Emoji
:smile: :rocket: :heart:

# Tables (simplified)
Header | Header
------ | ------
Cell   | Cell

# Alerts (GitHub)
> [!NOTE]
> Useful information

> [!WARNING]
> Critical content

> [!IMPORTANT]
> Key information
```
