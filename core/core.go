package core

type LinkIssueCommand struct {
	ProjectID string
	ID        int
}

type IMergeService interface {
	LinkIssue(cmd LinkIssueCommand) error
}

type IMergeStore interface {
	GetMerge(projectId string, id int) (MergeRequest, error)
	SaveDescription(merge MergeRequest) error
}

type MergeService struct {
	MergeStore IMergeStore
}

func (s MergeService) LinkIssue(cmd LinkIssueCommand) error {

	merge, err := s.MergeStore.GetMerge(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	merge.LinkIssue()

	s.MergeStore.SaveDescription(merge)

	return nil
}
