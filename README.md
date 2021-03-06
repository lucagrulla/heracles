# Heracles

Your favourite importer for your Withings scale.

Are you looking to move your Fitbit Aria scale data into Withings Health Mate and you are not sure how to convert files from one format to another?

Fear no more!
Heracles comes to rescue.

<div style="text-align: center">
    <img src="https://github.com/lucagrulla/heracles/blob/main/heracles.gif" alt="drawing" style="width:200px;"/>
</div>

## Installation

### Mac OSX

#### using [Homebrew](https://brew.sh)

```bash
brew tap lucagrulla/tap
brew install heracles
```

### Linux

#### using [Linuxbrew](https://linuxbrew.sh/brew/)

```bash
brew tap lucagrulla/tap
brew install heracles
```

#### .deb/.rpm

Download the ```.deb``` or ```.rpm``` from the [releases page](https://github.com/lucagrulla/heracles) and install with ````dpkg -i```` and ````rpm -i```` respectively.

Heracles will transform the fitbit export data in the Withings format.

## 3 Easy steps
#### 1 - Export Fitbit data
Export your data from Fitbit. Just follow the instruction [here](https://help.fitbit.com/articles/en_US/Help_article/1133.htm)
Unzip the file and locate your weight data.
They should be in a folder looking like this:
```
./MyFitbitData/<YourName>/Personal & Account
```
Where your name is your actual name in the Fitbit account.

### 2 - Use Heracles to do the magic
Run Heracles.
```console
~ $ heracles full_path_to_weight_data
```
Heracles will produce a number of csv files called `heracles.csv`, ready to be uploaded in Health Mate.

### 3 - Import your data into Health Mate
Import all the csv files created by Heracles into Health Mate following the instruction [here](https://support.withings.com/hc/en-us/articles/201491477-Health-Mate-Online-Dashboard-Importing-data)
