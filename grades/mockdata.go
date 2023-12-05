package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Week 1 Homework",
					Type:  GradeHomework,
					Score: 94,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 65,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Jane",
			LastName:  "Doe",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 77,
				},
				{
					Title: "Week 1 Homework",
					Type:  GradeHomework,
					Score: 0,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 65,
				},
			},
		},
	}
}
