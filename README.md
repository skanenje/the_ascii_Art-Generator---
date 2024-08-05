# Ascii Art Web Stylize
## Description
This project is a web based ascii art-Generator implemented in Go. It takes user input, generates ascii art based on selected banner, and displays the results on a web interface.

### Features
- ASII art generation based on user input.
- Web-based interface for easy interaction
- Error handling for invalid input and missing template files
- ASCII art output displayed using HTML and CSS.

### Authors

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/swabri-musa-565350291?lipi=urn%3Ali%3Apage%3Ad_flagship3_profile_view_base_contact_details%3Buf0Ls4oWR2O2WLUMO5sIBg%3D%3D>
            <img src=https://learn.zone01kisumu.ke/git/avatars/bc7899a0aac2630a0a9b50bf330437a7?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=skanenje/>
            <br />
            <sub style="font-size:14px"><b>skanenje</b></sub>
        </a>
    </td>
   
</tr>
</table>

## Usage
### How to Install
Ascii-art requires [go](https://go.dev/dl/)  v 1.22.2 to run

Once go is installed  clone into this repository to do this 

- Open terminal and run
``` sh
git clone https://github.com/skanenje/ascii-art-web-stylize
cd ascii-art-web-stylize
```
Once you are in the directory,

To run the program, use the following command, on the terminal command line:
```bash
go run main.go
```
The terminal will prompt 
```
Starting server at :8080
to access visit http://localhost:8080
```
Use the link provided to the access the web GUI.

1. **Access the Web Interface:**

    Open your web browser and navigate to the provided URL where the HTML file is hosted (i.e, http://localhost:8080/).

2. **Input Text:**

    You'll see a text area labeled "Enter Input Text Here:". Click inside this area and type the text you want to convert into ASCII art.

3. **Select a Banner Style:**

    Choose a banner style from the dropdown menu labeled "Select Banner:". Options include:
        Standard: Basic ASCII art representation.
        Thinkertoy: A more elaborate ASCII art style.
        Shadow: ASCII art with a shadow effect.

4. **Generate ASCII Art:**

    Click on the "Generate ASCII Art" button below the input form. This will process your input and display the ASCII art output below the form.

5. **View ASCII Art Output:**

    Once generated, the ASCII art output will appear under the heading "ASCII Art Output:". It will be displayed inside a bordered box for easy viewing.

7. **Reset Input:**

    To clear the input and start over, you can click on the "Reset" button located below the input form. This will refresh the page, clearing the text area and resetting the banner selection.

8. **Interact with the Output:**

    You can scroll within the ASCII art box to view larger outputs and adjust the view as needed.
## Implementation: Algorithms
- The program uses Go's standard packages for HTTP handling, template rendering, and file operations.
- It validates user input for ASCII characters only and checks for the existence and integrity of ASCII art templates.
- The web interface allows users to input text, and see the generated ASCII art.


## HOW TO CONTRIBUTE ? 👷 

**1.** Fork [this](https://github.com/skanenje/ascii-art-web-stylize) repository.

**2.** Clone the forked repository.

```terminal
git clone https://github.com/skanenje/ascii-art-web-stylize
```

**3.** Navigate to the project directory.

```terminal
cd ascii-art-web-stylize
```

**4.**  MAKE A NEW FOLDER WITH YOUR PROJECT NAME INSIDE 
<br>

**5.**  Also Add a README file in your project folder which consists of Description/screenshots about your project !
          
 
<br>

**6.** Create a new branch.

```terminal
git checkout -b <your_branch_name>
```

**7.** Add & Commit your changes.

```terminal
  git add .
  git commit -m "<your_commit_message>"
```

**7.** Push your local branch to the remote repository.

```terminal
git push -u origin <your_branch_name>
```

**8.** Create a Pull Request!

**Congratulations!** Sit and relax till we review your PR, you've made your contribution to (https://github.com/skanenje/ascii-art-web-stylize) project

### Issues and Contributions

If you encounter any issues or have suggestions for improvements, feel free to submit an issue or propose a change.

## Enjoy ASCII Art Web in Go!

Feel free to explore the codebase, run the program with different inputs, and contribute to enhancing the project. Happy coding experience!