package services

import (
	"github.com/enrico-laboratory/music-projects-database-notionapi/internal/models/parsedmodels"
	"github.com/enrico-laboratory/music-projects-database-notionapi/internal/models/unparsedmodels"
	"strings"
)

func parseMusicProject(u *unparsedmodels.MusicProject, p *parsedmodels.MusicProject) {

	var description string

	if len(u.Properties.Description.RichText) == 0 {
		description = ""
	} else {
		description = u.Properties.Description.RichText[0].PlainText
	}

	var choirRollup string

	if len(u.Properties.ChoirRollup.Rollup.Array) == 0 {
		choirRollup = ""
	} else {
		choirRollup = u.Properties.ChoirRollup.Rollup.Array[0].Title[0].PlainText
	}
	p.Id = u.ID
	p.CreatedTime = u.CreatedTime
	p.LastEditedTime = u.LastEditedTime
	p.Title = u.Properties.Title.Title[0].PlainText
	p.Year = u.Properties.Year.Number
	p.Status = u.Properties.Status.Select.Name
	p.Description = description
	p.ChoirRollup = choirRollup
}

func parseLocation(u *unparsedmodels.Location, p *parsedmodels.Location) {
	var contact string

	if len(u.Properties.Contact.Relation) == 0 {
		contact = ""
	} else {
		contact = u.Properties.Contact.Relation[0].ID
	}

	var phone string

	if len(u.Properties.Phone.Rollup.Array) == 0 {
		phone = ""
	} else {
		phone = u.Properties.Phone.Rollup.Array[0].PhoneNumber
	}

	var city string

	if len(u.Properties.City.RichText) == 0 {
		city = ""
	} else {
		city = u.Properties.City.RichText[0].PlainText
	}

	var email string

	if len(u.Properties.Email.Rollup.Array) == 0 {
		email = ""
	} else {
		email = u.Properties.Email.Rollup.Array[0].Email
	}

	var tasks []string

	if len(u.Properties.Tasks.Relation) == 0 {
		tasks = []string{}
	} else {
		for _, record := range u.Properties.Tasks.Relation {
			tasks = append(tasks, record.ID)
		}
	}

	var purpose []string

	if len(u.Properties.Purpose.MultiSelect) == 0 {
		purpose = []string{}
	} else {
		for _, record := range u.Properties.Purpose.MultiSelect {
			purpose = append(tasks, record.Name)
		}
	}

	var address string

	if len(u.Properties.Address.RichText) == 0 {
		address = ""
	} else {
		address = u.Properties.Address.RichText[0].PlainText
	}

	var location string

	if len(u.Properties.Location.Title) == 0 {
		location = ""
	} else {
		location = u.Properties.Location.Title[0].PlainText
	}

	p.Id = u.ID
	p.CreatedTime = u.CreatedTime
	p.LastEditedTime = u.LastEditedTime
	p.Contact = contact
	p.Phone = phone
	p.City = city
	p.Email = email
	p.Task = tasks
	p.Purpose = purpose
	p.Address = address
	p.Location = location

}

func parseContact(u *unparsedmodels.Contact, p *parsedmodels.Contact) {
	var singer string

	if len(u.Properties.SingerRollup.Rollup.Array) == 0 {
		singer = ""
	} else if len(u.Properties.SingerRollup.Rollup.Array[0].Title) == 0 {
		singer = ""
	} else {
		singer = u.Properties.SingerRollup.Rollup.Array[0].Title[0].PlainText
	}

	var singerId string

	if len(u.Properties.Singer.Relation) == 0 {
		singerId = ""
	} else {
		singerId = u.Properties.Singer.Relation[0].ID
	}

	var note string

	if len(u.Properties.Note.RichText) == 0 {
		note = ""
	} else {
		note = u.Properties.Note.RichText[0].PlainText
	}

	var phone string

	if len(u.Properties.Phone.Rollup.Array) == 0 {
		phone = ""
	} else {
		phone = u.Properties.Phone.Rollup.Array[0].PhoneNumber
	}

	var status string

	status = u.Properties.Status.Select.Name

	var order float64
	order = u.Properties.Order.Number

	var email string

	if len(u.Properties.Email.Rollup.Array) == 0 {
		email = ""
	} else {
		email = u.Properties.Email.Rollup.Array[0].Email
	}

	var musicProject []string

	if len(u.Properties.MusicProject.Relation) == 0 {
		musicProject = []string{}
	} else {
		for _, record := range u.Properties.MusicProject.Relation {
			musicProject = append(musicProject, record.ID)
		}
	}

	var role string

	if len(u.Properties.Role.Title) == 0 {
		role = ""
	} else {
		role = u.Properties.Role.Title[0].PlainText
	}

	p.Id = u.ID
	p.CreatedTime = u.CreatedTime
	p.LastEditedTime = u.LastEditedTime
	p.Singer = singer
	p.SingerId = singerId
	p.Note = note
	p.Phone = phone
	p.Status = status
	p.Order = order
	p.Email = email
	p.MusicProject = musicProject
	p.Role = role
}

func parsePiece(u *unparsedmodels.Piece, p *parsedmodels.Piece) {

	var order string

	if len(u.Properties.Order.Title) == 0 {
		order = ""
	} else {
		order = u.Properties.Order.Title[0].PlainText
	}

	var num1 string

	if len(u.Properties.OneTopVoice.RichText) == 0 {
		num1 = ""
	} else {
		num1 = u.Properties.OneTopVoice.RichText[0].PlainText
	}

	var num2 string

	if len(u.Properties.Num2.RichText) == 0 {
		num2 = ""
	} else {
		num2 = u.Properties.Num2.RichText[0].PlainText
	}

	var num3 string

	if len(u.Properties.Num3.RichText) == 0 {
		num3 = ""
	} else {
		num3 = u.Properties.Num3.RichText[0].PlainText
	}

	var num4 string

	if len(u.Properties.Num4.RichText) == 0 {
		num4 = ""
	} else {
		num4 = u.Properties.Num4.RichText[0].PlainText
	}

	var num5 string

	if len(u.Properties.Num5.RichText) == 0 {
		num5 = ""
	} else {
		num5 = u.Properties.Num5.RichText[0].PlainText
	}
	var num6 string

	if len(u.Properties.Num6.RichText) == 0 {
		num6 = ""
	} else {
		num6 = u.Properties.Num6.RichText[0].PlainText
	}

	var num7 string

	if len(u.Properties.Num7.RichText) == 0 {
		num7 = ""
	} else {
		num7 = u.Properties.Num7.RichText[0].PlainText
	}

	var num8 string

	if len(u.Properties.Num8.RichText) == 0 {
		num8 = ""
	} else {
		num8 = u.Properties.Num8.RichText[0].PlainText
	}

	var num9 string

	if len(u.Properties.Num9.RichText) == 0 {
		num9 = ""
	} else {
		num9 = u.Properties.Num9.RichText[0].PlainText
	}

	var num10 string

	if len(u.Properties.Num10.RichText) == 0 {
		num10 = ""
	} else {
		num10 = u.Properties.Num10.RichText[0].PlainText
	}

	var num11 string

	if len(u.Properties.Num11.RichText) == 0 {
		num11 = ""
	} else {
		num11 = u.Properties.Num11.RichText[0].PlainText
	}

	var num12 string

	if len(u.Properties.Num12.RichText) == 0 {
		num12 = ""
	} else {
		num12 = u.Properties.Num12.RichText[0].PlainText
	}

	var note string

	if len(u.Properties.Note.RichText) == 0 {
		note = ""
	} else {
		note = u.Properties.Note.RichText[0].PlainText
	}

	var voicing string

	if len(u.Properties.VoicesRollup.Rollup.Array) == 0 {
		voicing = ""
	} else {
		voicing = u.Properties.VoicesRollup.Rollup.Array[0].Select.Name
	}

	var solo string

	if len(u.Properties.SoloRollup.Rollup.Array) == 0 {
		solo = ""
	} else {
		solo = u.Properties.SoloRollup.Rollup.Array[0].Select.Name
	}

	var music string

	if len(u.Properties.MusicRollup.Rollup.Array) == 0 {
		music = ""
	} else {
		music = u.Properties.MusicRollup.Rollup.Array[0].Title[0].PlainText
	}

	var musicProjectsIds []string

	if len(u.Properties.MusicProject.Relation) == 0 {
		musicProjectsIds = append(musicProjectsIds, "")
	} else {
		for _, musicProjectsId := range u.Properties.MusicProject.Relation {
			musicProjectsIds = append(musicProjectsIds, musicProjectsId.ID)
		}
	}

	var media string

	if len(u.Properties.MediaRollup.Rollup.Array) == 0 {
		media = ""
	} else {
		media = u.Properties.MediaRollup.Rollup.Array[0].URL
	}

	var score string

	if len(u.Properties.ScoreRollup.Rollup.Array) == 0 {
		score = ""
	} else {
		score = u.Properties.ScoreRollup.Rollup.Array[0].URL
	}

	var recording string

	if len(u.Properties.RecordingRollup.Rollup.Array) == 0 {
		recording = ""
	} else {
		recording = u.Properties.RecordingRollup.Rollup.Array[0].URL
	}

	var instruments []string

	if len(u.Properties.InstrumentsRollup.Rollup.Array) == 0 {
		instruments = append(instruments, "")
	} else {
		for _, instrument := range u.Properties.InstrumentsRollup.Rollup.Array[0].MultiSelect {
			instruments = append(instruments, instrument.Name)
		}
	}

	var length float64

	if len(u.Properties.LengthRollup.Rollup.Array) == 0 {
		length = 0
	} else {
		length = u.Properties.LengthRollup.Rollup.Array[0].Number
	}

	var composer string

	if len(u.Properties.ComposerRollup.Rollup.Array) == 0 {
		composer = ""
	} else if len(u.Properties.ComposerRollup.Rollup.Array[0].RichText) == 0 {
		composer = ""
	} else {
		composer = strings.TrimSpace(u.Properties.ComposerRollup.Rollup.Array[0].RichText[0].PlainText)
	}

	p.Id = u.ID
	p.CreatedTime = u.CreatedTime
	p.LastEditedTime = u.LastEditedTime
	p.Order = order
	p.Num1 = num1
	p.Num2 = num2
	p.Num3 = num3
	p.Num4 = num4
	p.Num5 = num5
	p.Num6 = num6
	p.Num7 = num7
	p.Num8 = num8
	p.Num9 = num9
	p.Num10 = num10
	p.Num11 = num11
	p.Num12 = num12
	p.Note = note
	p.Voicing = voicing
	p.Solo = solo
	p.Music = music
	p.Selected = u.Properties.Selected.Checkbox
	p.MusicProject = musicProjectsIds
	p.Media = media
	p.Score = score
	p.Recording = recording
	p.Instrument = instruments
	p.Length = length
	p.Composer = composer

}

func parseTask(u *unparsedmodels.Task, p *parsedmodels.Task) {
	var musicProject []string

	if len(u.Properties.MusicProject.Relation) == 0 {
		musicProject = []string{}
	} else {
		for _, record := range u.Properties.MusicProject.Relation {
			musicProject = append(musicProject, record.ID)
		}
	}

	var _type string

	_type = u.Properties.Type.Select.Name

	var locationRollup string

	if len(u.Properties.LocationRollup.Rollup.Array) == 0 {
		locationRollup = ""
	} else if len(u.Properties.LocationRollup.Rollup.Array[0].Title) == 0 {
		locationRollup = ""
	} else {
		locationRollup = u.Properties.LocationRollup.Rollup.Array[0].Title[0].PlainText
	}

	var cityRollup string

	if len(u.Properties.LocationCityRollup.Rollup.Array) == 0 {
		cityRollup = ""
	} else if len(u.Properties.LocationCityRollup.Rollup.Array[0].RichText) == 0 {
		cityRollup = ""
	} else {
		cityRollup = u.Properties.LocationCityRollup.Rollup.Array[0].RichText[0].PlainText
	}

	var priority string

	priority = u.Properties.Priority.Select.Name

	var startDateAndTime string
	var endDateAndTime string
	emptyTaskDate := unparsedmodels.Date{}
	emptyDate := ""

	if u.Properties.DoDate.Date == emptyTaskDate {
		startDateAndTime = emptyDate
		endDateAndTime = emptyDate
	} else if u.Properties.DoDate.Date.Start == emptyDate &&
		u.Properties.DoDate.Date.End != emptyDate {
		startDateAndTime = emptyDate
		endDateAndTime = u.Properties.DoDate.Date.Start
	} else if u.Properties.DoDate.Date.Start != emptyDate &&
		u.Properties.DoDate.Date.End == emptyDate {
		startDateAndTime = u.Properties.DoDate.Date.Start
		endDateAndTime = emptyDate
	} else {
		startDateAndTime = u.Properties.DoDate.Date.Start
		endDateAndTime = u.Properties.DoDate.Date.Start
	}

	var duration string

	duration = u.Properties.Duration.Formula.String

	var kanban string
	emptySelect := unparsedmodels.Select{}

	if u.Properties.Kanban.Select == emptySelect {
		kanban = ""
	} else {
		kanban = u.Properties.Kanban.Select.Name
	}

	var isDone bool

	isDone = u.Properties.Done.Checkbox

	var locationId []string

	if len(u.Properties.Location.Relation) == 0 {
		locationId = []string{}
	} else {
		for _, record := range u.Properties.Location.Relation {
			locationId = append(locationId, record.ID)
		}
	}

	//var choirRollup string
	//
	//if len(u.Properties.ChoirRollup.Rollup.Array) == 0 {
	//	choirRollup = ""
	//} else if len(u.Properties.ChoirRollup.Rollup.Array[0].) == 0 {
	//	choirRollup = ""
	//} else {
	//	choirRollup = u.Properties.ChoirRollup.Rollup.Array[0].Title[0].PlainText
	//}

	var taskTitle string

	if len(u.Properties.Task.Title) == 0 {
		taskTitle = ""
	} else {
		taskTitle = u.Properties.Task.Title[0].PlainText
	}

	p.Id = u.ID
	p.CreatedTime = u.CreatedTime
	p.LastEditedTime = u.LastEditedTime
	p.MusicProject = musicProject
	p.Type = _type
	p.Location = locationRollup
	p.City = cityRollup
	p.Priority = priority
	p.StartDateAndTime = startDateAndTime
	p.EndDateAndTime = endDateAndTime
	p.Duration = duration
	p.Kanban = kanban
	p.IsDone = isDone
	p.LocationId = locationId
	//p.Choir= choirRollup,
	p.Title = taskTitle

}
