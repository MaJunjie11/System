package course_dal

import (
	"System/course/course_model"
	"System/db"
	"System/pb_gen"
	"System/user/user_moudle"
	"fmt"

	"github.com/jinzhu/gorm"
)

func TeacherAuditStudentToDb(courseId string, uid int64) error {
	db.Db.Model(&course_model.StudentSelectCourseInfo{}).Where("uid=? and course_id=?", uid, courseId).Update("status", 1)
	if err := db.Db.Model(&course_model.CourseBaseInfo{}).Where("course_id=?", courseId).Update("class_student_num", gorm.Expr("class_student_num+?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func GetStudentAddCourseListFromDbByUid(uid int64) *pb_gen.StudentGetAddCourseListData {

	// 多表联合查询

	var (
		result []course_model.CourseBaseInfo = make([]course_model.CourseBaseInfo, 0)
		count  int
		data   *pb_gen.StudentGetAddCourseListData = &pb_gen.StudentGetAddCourseListData{}
	)
	db.Db.Where("course_id not in (select course_id from student_select_course_infos where uid=?) and course_status=0", uid).Find(&result).Count(&count)
	if count < 1 {
		fmt.Println("no course to show")
		return nil
	}
	data.Count = int32(len(result))
	data.Rows = make([]*pb_gen.AddCourseData, 0)
	for _, item := range result {
		i := &pb_gen.AddCourseData{
			CourseId:        item.CourseId,
			TeacherName:     item.TeacherName,
			ClassName:       item.ClassName,
			ClassStudentNum: int32(item.ClassStudentNum),
			ClassCapacity:   int32(item.ClassCapacity),
		}
		data.Rows = append(data.Rows, i)
	}
	return data
}

func GetStudentNonJoinCourseByUid(uid int64) *pb_gen.StudentGetNonJoinCourseResponseData {
	var (
		result []course_model.CourseBaseInfo = make([]course_model.CourseBaseInfo, 0)
		count  int
		data   *pb_gen.StudentGetNonJoinCourseResponseData = &pb_gen.StudentGetNonJoinCourseResponseData{}
	)
	db.Db.Where("course_id not in (select course_id from student_select_course_infos where uid=?) and course_status!=0", uid).Find(&result).Count(&count)
	// 获取不能加入的课程
	if count < 1 {
		fmt.Println("no course to show")
		return nil
	}
	data.Total = int32(len(result))
	fmt.Println("data.total:", data.Total)
	data.Items = make([]*pb_gen.StudentNonJoinCourseData, 0)
	for _, item := range result {
		i := &pb_gen.StudentNonJoinCourseData{
			CourseId:        item.CourseId,
			TeacherName:     item.TeacherName,
			ClassName:       item.ClassName,
			ClassStudentNum: int32(item.ClassStudentNum),
			ClassCapacity:   int32(item.ClassCapacity),
		}
		if item.CourseStatus == 0 {
			i.ClassStatus = "未开课"
		} else if item.CourseStatus == 1 {
			i.ClassStatus = "开课"
		} else if item.CourseStatus == 2 {
			i.ClassStatus = "结课"
		}
		data.Items = append(data.Items, i)
	}
	return data
}

func GetAuditCourseStudentFromDbByCourseId(courseId string, uid int64) *pb_gen.GetAuditCourseStudentResponseData {
	var (
		result []*pb_gen.AuditCourseStudentData = make([]*pb_gen.AuditCourseStudentData, 0)
		count  int
		data   *pb_gen.GetAuditCourseStudentResponseData
	)
	db.Db.Table("user_student_base_info").Joins("left join student_select_course_infos on user_student_base_info.uid=student_select_course_infos.uid where student_select_course_infos.course_id=? and student_select_course_infos.status=0", courseId).Find(&result).Count(&count)
	if count < 1 {
		fmt.Println("no course to show")
		return nil
	}
	if len(result) == 0 {
		return nil
	}
	data = &pb_gen.GetAuditCourseStudentResponseData{}
	data.Total = int32(len(result))
	data.Items = result
	return data
}

func GetCourseStudentFromDbByCourseId(courseId string, uid int64) *pb_gen.GetCourseStudentResponseData {
	var (
		result []*pb_gen.CourseStudentData = make([]*pb_gen.CourseStudentData, 0)
		count  int
		data   *pb_gen.GetCourseStudentResponseData
	)
	db.Db.Table("user_student_base_info").Joins("left join student_select_course_infos on user_student_base_info.uid=student_select_course_infos.uid where student_select_course_infos.course_id=? and student_select_course_infos.status=1", courseId).Find(&result).Count(&count)
	if count < 1 {
		fmt.Println("no course to show")
		return nil
	}
	if len(result) == 0 {
		return nil
	}
	data = &pb_gen.GetCourseStudentResponseData{}
	data.Total = int32(len(result))
	data.Items = result
	return data
}

func GetStudentLikeCourseListFromDbByUid(uid int64) *pb_gen.StudentGetLikeCourseResponseData {
	var (
		result []course_model.CourseBaseInfo = make([]course_model.CourseBaseInfo, 0)
		count  int
		data   *pb_gen.StudentGetLikeCourseResponseData = &pb_gen.StudentGetLikeCourseResponseData{}
	)
	db.Db.Table("course_base_info").Select("*").Joins("inner join student_like_course_infos on course_base_info.course_id=student_like_course_infos.course_id where student_like_course_infos.uid = ?", uid).Scan(&result).Count(&count)
	if count < 1 {
		fmt.Println("no course to show")
		return nil
	}
	data.Total = int32(len(result))
	data.Items = make([]*pb_gen.StudentGetLikeCourseData, 0)
	for _, item := range result {
		i := &pb_gen.StudentGetLikeCourseData{
			CourseId:        item.CourseId,
			TeacherName:     item.TeacherName,
			ClassName:       item.ClassName,
			ClassStudentNum: int32(item.ClassStudentNum),
			ClassCapacity:   int32(item.ClassCapacity),
		}
		if item.CourseStatus == 0 {
			i.ClassStatus = "未开课"
		} else if item.CourseStatus == 1 {
			i.ClassStatus = "开课"
		} else if item.CourseStatus == 2 {
			i.ClassStatus = "结课"
		}
		data.Items = append(data.Items, i)
	}
	return data
}

func TeacherGetEndCouserFromDb(uid int64) *pb_gen.TeacherGetEndCourseResponseData {
	var (
		result []*pb_gen.TeacherCourseDetailData = make([]*pb_gen.TeacherCourseDetailData, 0)
		count  int
		data   *pb_gen.TeacherGetEndCourseResponseData = &pb_gen.TeacherGetEndCourseResponseData{}
	)

	db.Db.Table("course_base_info").Select("*").Joins("left join course_detail_info on course_base_info.course_id=course_detail_info.course_id where course_detail_info.teacher_uid = ? and course_detail_info.course_status=2", uid).Scan(&result).Count(&count)
	if count < 1 {
		fmt.Println("no course to show")
		return nil
	}
	if len(result) == 0 {
		return nil
	}

	data.Total = int32(len(result))
	data.Items = result
	return data
}

func TeacherGetStartCouserFromDb(uid int64) *pb_gen.TeacherGetStartCourseResponseData {
	var (
		result []*pb_gen.TeacherCourseDetailData = make([]*pb_gen.TeacherCourseDetailData, 0)
		count  int
		data   *pb_gen.TeacherGetStartCourseResponseData = &pb_gen.TeacherGetStartCourseResponseData{}
	)

	db.Db.Table("course_base_info").Select("*").Joins("left join course_detail_info on course_base_info.course_id=course_detail_info.course_id where course_detail_info.teacher_uid = ? and course_detail_info.course_status=1", uid).Scan(&result).Count(&count)
	if count < 1 {
		fmt.Println("no course to show")
		return nil
	}
	if len(result) == 0 {
		return nil
	}

	data.Total = int32(len(result))
	data.Items = result
	return data
}

func TeacherGetUntaughtCouserFromDb(uid int64) *pb_gen.TeacherGetUntaughtCourseResponseData {
	var (
		result []*pb_gen.TeacherCourseDetailData = make([]*pb_gen.TeacherCourseDetailData, 0)
		count  int
		data   *pb_gen.TeacherGetUntaughtCourseResponseData = &pb_gen.TeacherGetUntaughtCourseResponseData{}
	)

	db.Db.Table("course_base_info").Select("*").Joins("left join course_detail_info on course_base_info.course_id=course_detail_info.course_id where course_detail_info.teacher_uid = ? and course_detail_info.course_status=0", uid).Scan(&result).Count(&count)
	if count < 1 {
		fmt.Println("no course to show")
		return nil
	}
	if len(result) == 0 {
		return nil
	}

	data.Total = int32(len(result))
	data.Items = result
	return data
}

func GetStudentEndCourseListFromDbByUid(uid int64) *pb_gen.StudentGetEndCourseResponseData {
	var (
		result []*pb_gen.StudentCourseData = make([]*pb_gen.StudentCourseData, 0)
		count  int
		data   *pb_gen.StudentGetEndCourseResponseData = &pb_gen.StudentGetEndCourseResponseData{}
	)
	db.Db.Table("course_base_info").Select("*").Joins("left join student_select_course_infos on course_base_info.course_id=student_select_course_infos.course_id ").Joins("left join course_detail_info on course_base_info.course_id=course_detail_info.course_id where student_select_course_infos.uid = ? and student_select_course_infos.status = 1 and course_detail_info.course_status=2", uid).Scan(&result).Count(&count)
	if count < 1 {
		fmt.Println("no course to show")
		return nil
	}
	if len(result) == 0 {
		return nil
	}
	data.Total = int32(len(result))
	data.Items = make([]*pb_gen.StudentCourseData, 0)
	data.Items = result
	return data
}

func GetStudentStartCourseListFromDbByUid(uid int64) *pb_gen.StudentGetStartCourseResponseData {
	var (
		result []*pb_gen.StudentCourseData = make([]*pb_gen.StudentCourseData, 0)
		count  int
		data   *pb_gen.StudentGetStartCourseResponseData = &pb_gen.StudentGetStartCourseResponseData{}
	)
	db.Db.Table("course_base_info").Select("*").Joins("left join student_select_course_infos on course_base_info.course_id=student_select_course_infos.course_id ").Joins("left join course_detail_info on course_base_info.course_id=course_detail_info.course_id where student_select_course_infos.uid = ? and student_select_course_infos.status = 1 and course_detail_info.course_status=1", uid).Scan(&result).Count(&count)
	if count < 1 {
		fmt.Println("no course to show")
		return nil
	}
	if len(result) == 0 {
		return nil
	}
	data.Total = int32(len(result))
	data.Items = make([]*pb_gen.StudentCourseData, 0)
	data.Items = result
	return data
}

func GetStudentUntaughtCourseListFromDbByUid(uid int64) *pb_gen.StudentGetUntaughtCourseResponseData {
	var (
		result []*pb_gen.StudentCourseData = make([]*pb_gen.StudentCourseData, 0)
		count  int
		data   *pb_gen.StudentGetUntaughtCourseResponseData = &pb_gen.StudentGetUntaughtCourseResponseData{}
	)
	db.Db.Table("course_base_info").Select("*").Joins("left join student_select_course_infos on course_base_info.course_id=student_select_course_infos.course_id ").Joins("left join course_detail_info on course_base_info.course_id=course_detail_info.course_id where student_select_course_infos.uid = ? and student_select_course_infos.status = 1 and course_detail_info.course_status=0", uid).Scan(&result).Count(&count)
	if count < 1 {
		fmt.Println("no course to show")
		return nil
	}
	if len(result) == 0 {
		return nil
	}
	data.Total = int32(len(result))
	data.Items = make([]*pb_gen.StudentCourseData, 0)
	data.Items = result
	return data
}

func TeacherStartCourse(courseId string) error {
	db.Db.Model(&course_model.CourseBaseInfo{}).Where("course_id=?", courseId).Update("course_status", 1)

	db.Db.Model(&course_model.CourseDetailInfo{}).Where("course_id=?", courseId).Update("course_status", 1)
	fmt.Println("开课成功")
	return nil
}

func TeacherDeleteRefuseCourse(courseId string) error {
	db.Db.Where("course_id=?", courseId).Delete(&course_model.CourseRefuseInfo{})
	return nil
}

func TeacherEndCourse(courseId string) error {
	db.Db.Model(&course_model.CourseBaseInfo{}).Where("course_id=?", courseId).Update("course_status", 2)
	db.Db.Model(&course_model.CourseDetailInfo{}).Where("course_id=?", courseId).Update("course_status", 2)
	return nil
}

func TeacherAddStudent(uuid string, course_id string) error {
	var (
		err   error
		data  course_model.StudentSelectCourseInfo
		count int
	)
	// TODO: 这里需要做学生校验
	userInfo := &user_moudle.UserStudentBaseInfo{}
	db.Db.Where("uuid = ?", uuid).Find(userInfo).Count(&count)
	if count < 1 {
		return fmt.Errorf("no student")
	}

	data.CourseId = course_id
	data.Uid = userInfo.Uid
	data.Status = 1 // 等待审核状态 status 1 通过审核的
	db.Db.Create(&data)

	// 学生数量+1
	if err = db.Db.Model(&course_model.CourseBaseInfo{}).Where("course_id=?", course_id).Update("class_student_num", gorm.Expr("class_student_num+?", 1)).Error; err != nil {
		return err
	}
	return err
}

func StudentAddCourse(uid int64, course_id string) error {
	var (
		err  error
		data course_model.StudentSelectCourseInfo
	)
	data.CourseId = course_id
	data.Uid = uid
	data.Status = 0 // 等待审核状态 status 1 通过审核的
	db.Db.Create(&data)
	return err
}

func StudentAddLikeCourse(uid int64, course_id string) error {
	var (
		err  error
		data course_model.StudentLikeCourseInfo
	)
	data.CourseId = course_id
	data.Uid = uid
	db.Db.Create(&data)
	return err
}

// 存入信息
func AddCourseAuditInfo(data *course_model.CourseAuditInfo) error {
	var (
		err error
	)
	db.Db.Create(data)
	return err
}

func TeacherGetRefuseInfoFromDb(uid int64) *pb_gen.TeacherGetRefuseCourseResponseData {
	var (
		data   *pb_gen.TeacherGetRefuseCourseResponseData = &pb_gen.TeacherGetRefuseCourseResponseData{}
		result []course_model.CourseRefuseInfo            = make([]course_model.CourseRefuseInfo, 0)
		count  int
	)
	db.Db.Where("teacher_uid = ?", uid).Find(&result).Count(&count)
	if count < 1 {
		fmt.Println("no course to show")
		return nil
	}
	// 组织data
	data.Total = int32(len(result))
	if data.Total == 0 {
		return nil
	}
	for _, item := range result {
		data1 := &pb_gen.RefuseCourseData{
			ClassName:       item.ClassName,
			CourseMajor:     item.CourseMajor,
			CourseDate:      item.CourseDate,
			ClassCapacity:   int32(item.ClassCapacity),
			CourseIntroduce: item.CourseIntroduce,
			CourseId:        item.CourseId,
			RefuseReason:    item.RefuseReason,
		}
		data.Items = append(data.Items, data1)
	}
	return data
}

func TeacherGetAuditInfoFromDb(uid int64) *pb_gen.TeacherGetAuditCourseResponseData {
	var (
		data   *pb_gen.TeacherGetAuditCourseResponseData = &pb_gen.TeacherGetAuditCourseResponseData{}
		result []course_model.CourseAuditInfo            = make([]course_model.CourseAuditInfo, 0)
		count  int
	)
	db.Db.Where("teacher_uid = ?", uid).Find(&result).Count(&count)
	if count < 1 {
		fmt.Println("no course to show")
		return nil
	}
	// 组织data
	data.Total = int32(len(result))
	if data.Total == 0 {
		return nil
	}
	for _, item := range result {
		data1 := &pb_gen.AuditCourseData{
			ClassName:       item.ClassName,
			CourseMajor:     item.CourseMajor,
			CourseDate:      item.CourseDate,
			ClassCapacity:   int32(item.ClassCapacity),
			CourseIntroduce: item.CourseIntroduce,
		}
		data.Items = append(data.Items, data1)
	}
	return data
}

func ManagerResetPassword(uuid string) error {
	db.Db.Model(&user_moudle.UserStudentLoginInfo{}).Where("uuid = ?", uuid).Update("password", "")
	return nil
}

func ManagerAddRefuseCourseToDb(uid int64, courseId string, refuseReason string) error {
	var (
		result *course_model.CourseAuditInfo = &course_model.CourseAuditInfo{}
		count1 int
		data   *course_model.CourseRefuseInfo
	)
	db.Db.Where("course_id=?", courseId).Find(result).Count(&count1)
	if count1 < 1 {
		fmt.Println("no course to show")
		return nil
	}

	data = &course_model.CourseRefuseInfo{
		CourseId:        result.CourseId,
		TeacherUid:      result.TeacherUid,
		ClassName:       result.ClassName,
		CourseDate:      result.CourseDate,
		ClassCapacity:   result.ClassCapacity,
		CourseMajor:     result.CourseMajor,
		CourseIntroduce: result.CourseIntroduce,
		ManagerUid:      uid,
		RefuseReason:    refuseReason,
	}
	db.Db.Create(data)

	db.Db.Where("course_id=?", courseId).Delete(&course_model.CourseAuditInfo{})
	return nil
}

// 根据managerID course_id 将审核列表数据加入正式的表中
func ManagerAuditDownCourseToDb(uid int64, courseId string, courseAddr string) error {
	var (
		result       *course_model.CourseAuditInfo    = &course_model.CourseAuditInfo{}
		userBaseInfo *user_moudle.UserStudentBaseInfo = &user_moudle.UserStudentBaseInfo{}
		count1       int
		count2       int
	)

	db.Db.Where("course_id=?", courseId).Find(result).Count(&count1)
	if count1 < 1 {
		fmt.Println("no course to show")
		return nil
	}
	db.Db.Where("uid=?", result.TeacherUid).Find(userBaseInfo).Count(&count2)
	if count2 < 1 {
		fmt.Println("no course to show")
		return nil
	}

	courseBaseInfo := course_model.CourseBaseInfo{
		CourseId:        result.CourseId,
		TeacherName:     userBaseInfo.Name,
		ClassName:       result.ClassName,
		ClassStudentNum: 0,
		ClassCapacity:   result.ClassCapacity,
		CourseStatus:    0, //未开课
	}
	courseDetailInfo := course_model.CourseDetailInfo{
		CourseId:        result.CourseId,
		TeacherUid:      userBaseInfo.Uid,
		TeacherUuid:     userBaseInfo.Uuid,
		CourseAddr:      courseAddr,
		CourseDate:      result.CourseDate,
		CourseIntroduce: result.CourseIntroduce,
		CourseMajor:     result.CourseMajor,
		CourseStatus:    0,
		ManagerUid:      uid,
	}

	db.Db.Create(&courseBaseInfo)
	db.Db.Create(&courseDetailInfo)

	// 删除这个数据
	db.Db.Where("course_id=?", courseId).Delete(&course_model.CourseAuditInfo{})
	return nil
}

func FeetStudentInDb(uid int64, courseId string) error {
	var (
		data *course_model.StudentSelectCourseInfo = &course_model.StudentSelectCourseInfo{}
	)

	db.Db.Where("course_id=? and uid = ?", courseId, uid).Find(data)
	if data.Status == 1 {
		db.Db.Model(&course_model.CourseBaseInfo{}).Where("course_id=?", courseId).Update("class_student_num", gorm.Expr("class_student_num-?", 1))
	}
	db.Db.Where("course_id=? and uid=?", courseId, uid).Delete(course_model.StudentSelectCourseInfo{})
	return nil
}

// 根据managerID拿到管理员需要审核的课程列表
func ManagerGetAuditCourseFromDb(manager_uid int64) *pb_gen.ManagerGetAuditCourseResponseData {
	var (
		data   *pb_gen.ManagerGetAuditCourseResponseData = &pb_gen.ManagerGetAuditCourseResponseData{}
		result []*pb_gen.AuditCourseData                 = make([]*pb_gen.AuditCourseData, 0)
		count  int
	)

	db.Db.Table("course_audit_infos").Select("*").Joins("left join user_student_base_info on course_audit_infos.teacher_uid=user_student_base_info.uid where course_audit_infos.manager_uid = ?", manager_uid).Find(&result).Count(&count)
	if count < 1 {
		fmt.Println("no course to show")
		return nil
	}
	// 组织data
	data.Total = int32(len(result))
	if data.Total == 0 {
		return nil
	}
	data.Items = result
	fmt.Println("course_id:", data.Items[0].CourseId)
	return data
}

func GetManagerUidFromUuid(uuid string) int64 {
	var (
		count int
	)
	userInfo := &user_moudle.UserStudentBaseInfo{}
	db.Db.Where("uuid = ?", uuid).Find(userInfo).Count(&count)
	if count < 1 {
		fmt.Println("no course to show")
		return -1
	}
	if userInfo.Limit != 3 {
		fmt.Println("不是管理员")
		return -1
	}
	return userInfo.Uid
}
