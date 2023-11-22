# Excel (Spreadsheet) to JSON

## Purpose
Convert the Report Recommendations data from its source Excel spreadsheet into JSON-string.

## Running The App

### Docker Compose
```bash
$ git clone git@gitlab.oit.duke.edu:dlc32/excel-to-json.git
$ cd /path/to/excel-to-json
$ docker-compose up --build

# enjoy the show.
```
### Locally Native
You'll need at least version `1.17` of `golang`.  
  
```bash
$ git clone git@gitlab.oit.duke.edu:dlc32/excel-to-json.git
$ cd /path/to/excel-to-json
$ go run .
```
