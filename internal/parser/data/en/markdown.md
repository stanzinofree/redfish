# Markdown Syntax Guide

## markdown headers
**Tags**: markdown, headers, headings, title, h1, h2, h3
**Keywords**: headers headings title h1 h2 h3 h4 h5 h6
**Short_Description**: Create headers
**Long_Description**: Defines document structure with six heading levels using hash symbols. H1 is the largest and most important, H6 is the smallest. Number of hashes (1-6) determines heading level. Essential for organizing content hierarchy and improving document readability.

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
**Short_Description**: Text formatting
**Long_Description**: Applies text styling using asterisks or underscores. Single asterisk/underscore for italic, double for bold, triple for both. Tildes create strikethrough, backticks create inline code. Combines for rich text formatting without HTML.

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
**Short_Description**: Create lists
**Long_Description**: Creates structured lists using hyphens/asterisks for bullets or numbers for ordered lists. Supports nesting with indentation. Task lists use brackets with x for completed items. Essential for organizing information and creating checklists.

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
**Short_Description**: Create links
**Long_Description**: Creates clickable hyperlinks with text and URL. Supports inline links, reference-style links for reuse, and auto-links for plain URLs. Optional title attribute appears on hover. Reference style keeps document clean when same link appears multiple times.

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
**Short_Description**: Embed images
**Long_Description**: Embeds images using syntax similar to links with preceding exclamation mark. Alt text describes image for accessibility. Supports inline and reference-style syntax. Can wrap in link to make image clickable. Title attribute optional.

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
**Short_Description**: Code blocks with syntax highlighting
**Long_Description**: Creates fenced code blocks using triple backticks with optional language identifier for syntax highlighting. Supports hundreds of languages including JavaScript, Python, Go, Bash, etc. Essential for technical documentation and sharing code snippets.

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
**Short_Description**: Create blockquotes
**Long_Description**: Creates quoted text blocks using greater-than symbol. Supports multi-line quotes, multiple paragraphs, and nesting with additional symbols. Commonly used for citations, callouts, or emphasizing important information. Can contain other markdown formatting.

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
**Short_Description**: Create tables
**Long_Description**: Creates data tables using pipes and hyphens. Header row separated by alignment row with hyphens. Colons control alignment: left (default), center (:---:), or right (---:). Essential for presenting structured data in readable format.

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
**Short_Description**: Create horizontal rules
**Long_Description**: Creates visual separators using three or more hyphens, asterisks, or underscores. Creates horizontal line spanning document width. Useful for breaking up sections, separating content, or creating visual hierarchy in long documents.

```markdown
---
***
___
```

## markdown footnotes
**Tags**: markdown, footnotes, references, citations, notes
**Keywords**: footnotes references citations notes superscript
**Short_Description**: Add footnotes
**Long_Description**: Creates footnotes with numbered references in text and definitions at bottom. Uses bracket-caret-number syntax for reference, same syntax with colon for definition. Supports multiple references to same footnote. Useful for citations and additional context.

```markdown
Here's a sentence with a footnote[^1].

[^1]: This is the footnote content.

Multiple references[^note] to the same note[^note].

[^note]: Shared footnote.
```

## markdown definition lists
**Tags**: markdown, definition, lists, terms, glossary
**Keywords**: definition lists terms glossary dictionary
**Short_Description**: Create definition lists
**Long_Description**: Creates glossary-style lists with terms and definitions. Term on one line, definition on next line starting with colon and space. Supports multiple definitions per term. Ideal for glossaries, documentation, and FAQs.

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
**Short_Description**: Embed raw HTML
**Long_Description**: Allows embedding raw HTML for features not available in Markdown. Supports inline and block-level HTML elements. Useful for custom styling, interactive elements like details/summary, or keyboard shortcuts. Most Markdown processors allow HTML passthrough.

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
**Short_Description**: Escape special characters
**Long_Description**: Uses backslash to display special Markdown characters literally without triggering formatting. Essential when you need to show actual asterisks, brackets, or other syntax characters. Allows writing about Markdown syntax in Markdown documents.

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
**Short_Description**: GitHub Flavored Markdown features
**Long_Description**: GitHub-specific Markdown extensions including task lists, @mentions, issue references (#123), emoji shortcodes, and alert callouts. GFM also includes table support and automatic URL linking. These features work on GitHub but may not work in all Markdown processors.

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
