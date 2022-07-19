package notion

import "github.com/enrico-laboratory/music-projects-database-notionapi/internal/models/unparsedmodels"

type Service interface {
	GetMusicProjects(body string, musicProject *unparsedmodels.MusicProjects) error
	GetMusicProjectsByStatus(status string, musicProject *unparsedmodels.MusicProjects) error
	GetRepertoireByProjectId(projectId string, musicProject *unparsedmodels.Repertoire) error
	GetCastByProjectId(projectId string, musicProject *unparsedmodels.Cast) error
	GetScheduleByProjectId(projectId string, musicProject *unparsedmodels.Schedule) error

	GetMusicProjectByID(id string, musicProject *unparsedmodels.MusicProject) error
	GetLocationById(id string, location *unparsedmodels.Location) error
}
