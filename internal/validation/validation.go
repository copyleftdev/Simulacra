package validation

import (
    "fmt"

    "simulacra/pkg/types"
)

func ValidatePlaybook(playbook *types.Playbook) error {
    for _, action := range playbook.Actions {
        if action.Type == "" {
            return fmt.Errorf("action type is required")
        }
        switch action.Type {
        case "mouse_move":
            if action.X < 0 || action.Y < 0 {
                return fmt.Errorf("invalid mouse coordinates")
            }
        case "key_press":
            if action.Key == "" {
                return fmt.Errorf("key is required for key_press action")
            }
        case "network_request":
            if action.URL == "" || action.Method == "" {
                return fmt.Errorf("url and method are required for network_request action")
            }
        default:
            return fmt.Errorf("unknown action type: %s", action.Type)
        }
    }
    return nil
}
