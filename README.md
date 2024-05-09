# Archi Git hooks

## Overview
This repository contains Git commit-msg hook designed to work with the Archi modeling toolkit and coArchi plugin for model collaboration. 
Commit-msg hook enriches Archi commit messages with the list of changed Archimate diagrams (views) in current commit. 

**Archi diagram name is used.** It is very usefull to additionally check your model changes operating not only Archi (and coArchi) identifiers from `git status` or Merge Request files list. 

**Changes that will be captured:**
- Any geometry changes in the diagram (moving, deletion any element from view etc.)
- Any changes to the diagram name, documentation, properties etc.
- Any changes with visual Archi (not Archimate!) objects (Note, Group, Connection) in the diagram
- Diagram deletion

**Changes that WILL NOT BE CAPTURED:**
- Any changes to the element (or relation) name, documentation, properties etc. even if this element is used in a diagram

### Without this hook
```
commit de6fb0573cdc5051d8b213de2dc31da615c1e298
Author: Denis Mazilov
Date:   Thu May 9 18:55:48 2024 +0300

    Added very important things to my Archi model!

```

### With this hook
```
commit dd6fb0573cdc5051d8b213de2dc31da615c1e298
Author: Denis Mazilov
Date:   Thu May 9 18:55:48 2024 +0300

    Added very important things to my Archi model!
    
    Following Archimate diagrams were changed in current commit:
    
    1) Best diagram ever [ id-c4bef9565657445c1ad7cc3b132ef238 ]
    2) My first Archi diagram [ id-a7c627b1ac4842e38716d0d0f323e0e2 ]
    3) DELETED diagram [ model/diagrams/id-c6301799364c42fea7a22f4123b3ae6e/ArchimateDiagramModel_id-26752c01590e4f6282b75df99b9f42bd.xml ]

```


## Getting Started

### Prerequisites

To use the commit-msg hook from this repo, you must have the following installed:
- Git https://git-scm.com/downloads/
- Archi modeling toolkit https://www.archimatetool.com/download/
- coArchi plugin https://www.archimatetool.com/plugins/

Optionally, if you want to build the binary yourself, you'll need to install Go. https://golang.org/dl/ 

### Default steps

1. **Clone the repository**:
   ```sh
   git clone https://github.com/dmazilov/archi-git-hooks.git
   ```
2. **Enter project directory**:
    ```sh   
   cd archi-git-hooks
   ``` 

### Prebuilded binary
For case you don't want install Go and build script yourself I prepared binary file.
To use it you must follow these steps:

3. **Make your hook file executable**:
    ```sh   
   chmod +x bin/commit-msg
   ```

4. **Copy prepared hook to your repository**:
    ```sh   
   cp bin/commit-msg your/archi/model/repo/.git/hooks/
   
### Self building
If you want to build git hook binary youself follow these steps:
   
3. **Build Go script**:
    ```sh   
   go build -o commit-msg commit_msg.go
   ```

4. **Make your hook file executable**:
    ```sh   
   chmod +x commit-msg
   ```

5. **Copy prepared hook to your repository**:
    ```sh   
   cp commit-msg your/archi/model/repo/.git/hooks/
   ```

## License
This project is licensed under the MIT License - see the LICENSE.md file for details.