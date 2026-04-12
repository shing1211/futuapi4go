name: Feature Request
description: Suggest a new feature or improvement
title: "[Feature] "
labels: ["enhancement"]
assignees: []
body:
  - type: markdown
    id: description
    attributes:
      value: |
        ## Feature Description
        <!-- Describe the feature or improvement you want. Be specific. -->
  - type: markdown
    id: usecase
    attributes:
      value: |
        ## Use Case
        <!-- Why do you need this feature? What's the problem you're solving? -->
  - type: markdown
    id: solution
    attributes:
      value: |
        ## Proposed Solution
        <!-- Describe your proposed solution or implementation approach -->
  - type: markdown
    id: alternatives
    attributes:
      value: |
        ## Alternatives Considered
        <!-- Describe any alternative solutions you've considered -->
  - type: textarea
    id: mockups
    attributes:
      label: Mockups / Examples
      description: Any code examples, mockups, or diagrams
      placeholder: |
        ```go
        // Example usage
        ```
  - type: checkboxes
    id: checklist
    attributes:
      label: Checklist
      options:
        - label: "I have searched existing issues and confirmed this is a new feature"
          required: true
        - label: "I am willing to help implement this feature"
          required: false
        - label: "This feature does not change the public API in a breaking way"
          required: false
