# kraken1
C2 Framework.

## Installation

Clone Project
```
git clone https://github.com/billx00/kraken1.git
```

Go into Project's directory
```
cd kraken1/
```

Run the C2-Framework
```
go run .
```

## Usage

Generate a payload with the generate command
Example:
```
>> generate os=linux l=127.0.0.1:8080
```

Display Sessions
```
>> sessions
```

Run a shell against a session
```
>> shell {SESSION-ID}
```

### Run Exploits (STILL IN TEST NOT REAlly USEABLE)

List all exploits
```
>> exploits
```

Select exploit
```
use exploit/linux/eurl
```

See all exploit options
```
>> options
```

Set exploit options
Example:
```
>> set option1 value
```
```
>> set sessions {SESSION-ID}
```

Run exploit
```
exploit
```
