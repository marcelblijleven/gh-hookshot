package tabs

func getNextTab(tabLen, current int) int {
	if current+1 == tabLen {
		return 0
	}

	return current + 1
}

func getPreviousTab(tabLen, current int) int {
	if current-1 < 0 {
		return tabLen - 1
	}

	return current - 1
}

func renderTab(name string, activeIdx int) string {
	idx := tabMap[name]

	if idx == activeIdx {
		return activeTab.Render(name)
	}

	return tab.Render(name)
}
