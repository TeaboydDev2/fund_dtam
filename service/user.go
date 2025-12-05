package service

// func (us *UserService) CreateUser(ctx context.Context, user *entities.User, file *entities.FileObject, illustration []*entities.FileObject) error {

// 	filePath := fmt.Sprintf("profile_picture/%s", uuid.New().String())

// 	if err := us.fileRepository.Upload(ctx, filePath, file.ContentType, file.File, file.Size); err != nil {
// 		return err
// 	}

// 	user.ProfilePicture.Ext = file.Ext
// 	user.ProfilePicture.Alt = file.Alt
// 	user.ProfilePicture.Path = filePath

// 	user.Illustration = make([]*entities.FileObject, len(illustration))

// 	for i, v := range illustration {

// 		multiFilePath := fmt.Sprintf("illustration_picture/%s", uuid.New().String())

// 		if err := us.fileRepository.Upload(ctx, multiFilePath, v.ContentType, v.File, v.Size); err != nil {
// 			return err
// 		}

// 		user.Illustration[i] = &entities.FileObject{
// 			Alt:  v.Alt,
// 			Ext:  v.Ext,
// 			Path: multiFilePath,
// 		}
// 	}

// 	return us.userRepository.SaveUser(ctx, user)
// }
