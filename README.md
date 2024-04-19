# qualitinvest
my quality investing journey

# WORKING WITH THE PROJECT
This section gives advices on how to run this project Python Flask project

## Virtual environments
What problem does a virtual environment solve? The more Python projects you have, the more likely it is that you need to work with different versions of Python libraries, or even Python itself. Newer versions of libraries for one project can break compatibility in another project.

Virtual environments are independent groups of Python libraries, one for each project. Packages installed for one project will not affect other projects or the operating system’s packages.

Python comes bundled with the venv module to create virtual environments.

Create a project folder and a .venv folder within:
```bash
python3.11 -m venv .venv
```

Before you work on your project, activate the corresponding environment:
```bash
. .venv/bin/activate 
```

Within the activated environment, use the following command to install the project dependencies:
```bash
pip install -r requirements.txt 
```

## Run The Application
Once the project dependencies are installed, you can run the Flask application running the following command:

```bash
flask --app app run
```

### Debug Mode
You might want to run the Flask application with the debug mode actived. By enabling debug mode, the server will automatically reload if code changes, and will show an interactive debugger in the browser if an error occurs during a request.

To enable debug mode, use the `--debugger` option:

```bash
flask --app app run --debugger 
```