name: Bug Report
description: Report something that is broken or not working as expected
title: "[Bug] "
labels: ["bug"]
assignees: []
body:
  - type: markdown
    id: description
    attributes:
      value: |
        ## Bug Description
        <!-- Describe the bug clearly and concisely. What went wrong? -->
  - type: textarea
    id: steps
    attributes:
      label: Steps to Reproduce
      description: How can we reproduce this? (Be as specific as possible)
      placeholder: |
        1. Go to '...'
        2. Call '...'
        3. See error
    validations:
      required: true
  - type: markdown
    id: expected
    attributes:
      value: |
        ## Expected Behavior
        <!-- What should have happened? -->
  - type: markdown
    id: actual
    attributes:
      value: |
        ## Actual Behavior
        <!-- What actually happened? Include error messages, stack traces, etc. -->
  - type: markdown
    id: environment
    attributes:
      value: |
        ## Environment
        <!-- Fill in your environment details -->
        | Item | Value |
        |------|-------|
        | **Go Version** | `go version` |
        | **futuapi4go Version** | e.g., v0.6.0 |
        | **OS / Arch** | e.g., Linux amd64 |
        | **Futu OpenD Version** | e.g., v6.x |
        | **Connection** | TCP / WebSocket |
  - type: textarea
    id: logs
    attributes:
      label: Logs / Code
      description: Relevant logs, code snippets, or stack traces
      placeholder: |
        ```
        paste logs or code here
        ```
  - type: checkboxes
    id: checklist
    attributes:
      label: Checklist
      options:
        - label: "I have searched existing issues and confirmed this is a new bug"
          required: true
        - label: "I can reproduce the issue with a minimal example"
          required: true
        - label: "The issue persists after restarting Futu OpenD"
          required: false
