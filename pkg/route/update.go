package route

func update(updater Updater) error {
	routes, err := List()
	if err != nil {
		return err
	}

	routes = updater(routes)

	if err := write(routes); err != nil {
		return err
	}

	return nil
}
