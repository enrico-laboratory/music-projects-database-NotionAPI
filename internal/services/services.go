package services

import (
	"github.com/enrico-laboratory/music-projects-database-notionapi/internal/models/parsedmodels"
	"github.com/enrico-laboratory/music-projects-database-notionapi/internal/models/unparsedmodels"
	"github.com/enrico-laboratory/music-projects-database-notionapi/internal/notion"
)

type Service struct {
	notionClient *notion.Client
}

func NewService(c *notion.Client) Service {
	return Service{
		notionClient: c,
	}
}

func (s *Service) GetMusicProjects(model *[]parsedmodels.MusicProject) error {

	var musicProjects unparsedmodels.MusicProjects
	err := s.notionClient.GetMusicProjects("", &musicProjects)
	if err != nil {
		return err
	}

	for _, musicProject := range musicProjects.Results {
		var musicProjectParsed parsedmodels.MusicProject
		parseMusicProject(&musicProject, &musicProjectParsed)
		*model = append(*model, musicProjectParsed)
	}
	return nil
}

func (s *Service) GetMusicProjectsByStatus(status string, model *[]parsedmodels.MusicProject) error {
	var musicProjects unparsedmodels.MusicProjects
	err := s.notionClient.GetMusicProjectsByStatus(status, &musicProjects)
	if err != nil {
		return err
	}

	for _, musicProject := range musicProjects.Results {
		var musicProjectParsed parsedmodels.MusicProject
		parseMusicProject(&musicProject, &musicProjectParsed)
		*model = append(*model, musicProjectParsed)
	}

	return nil
}

func (s *Service) GetMusicProjectById(id string, model *parsedmodels.MusicProject) error {
	var musicProject unparsedmodels.MusicProject
	err := s.notionClient.GetMusicProjectByID(id, &musicProject)
	if err != nil {
		return err
	}
	parseMusicProject(&musicProject, model)

	return nil
}

func (s *Service) GetRepertoireByProjectId(projectId string, model *[]parsedmodels.Piece) error {
	var repertoire unparsedmodels.Repertoire
	err := s.notionClient.GetRepertoireByProjectId(projectId, &repertoire)
	if err != nil {
		return err
	}

	for _, piece := range repertoire.Results {
		var pieceParsed parsedmodels.Piece
		parsePiece(&piece, &pieceParsed)
		*model = append(*model, pieceParsed)
	}
	return nil
}

func (s *Service) GetCastByProjectId(projectId string, model *[]parsedmodels.Contact) error {
	var cast unparsedmodels.Cast
	err := s.notionClient.GetCastByProjectId(projectId, &cast)
	if err != nil {
		return err
	}

	for _, contact := range cast.Results {
		var contactParsed parsedmodels.Contact
		parseContact(&contact, &contactParsed)
		*model = append(*model, contactParsed)
	}
	return nil
}

func (s *Service) GetScheduleByProjectId(projectId string, model *[]parsedmodels.Task) error {
	var schedule unparsedmodels.Schedule
	err := s.notionClient.GetScheduleByProjectId(projectId, &schedule)
	if err != nil {
		return err
	}

	for _, task := range schedule.Results {
		var taskParsed parsedmodels.Task
		parseTask(&task, &taskParsed)
		*model = append(*model, taskParsed)
	}
	return nil
}

func (s *Service) GetLocationById(id string, model *parsedmodels.Location) error {
	var location unparsedmodels.Location
	err := s.notionClient.GetLocationById(id, &location)
	if err != nil {
		return err
	}
	parseLocation(&location, model)
	return nil
}
