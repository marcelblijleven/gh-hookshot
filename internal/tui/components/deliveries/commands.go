package deliveries

import tea "github.com/charmbracelet/bubbletea"

func deliverySelectedCmd(hookID, deliveryID int) tea.Cmd {
	return func() tea.Msg {
		return DeliverySelectedMsg{HookID: hookID, DeliveryID: deliveryID}
	}
}
