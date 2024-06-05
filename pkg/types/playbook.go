package types

type Action struct {
    Type   string `yaml:"type"`
    X      int    `yaml:"x,omitempty"`
    Y      int    `yaml:"y,omitempty"`
    Key    string `yaml:"key,omitempty"`
    URL    string `yaml:"url,omitempty"`
    Method string `yaml:"method,omitempty"`
    Delay  int    `yaml:"delay"`
}

type Playbook struct {
    Name    string   `yaml:"name"`
    Actions []Action `yaml:"actions"`
}
