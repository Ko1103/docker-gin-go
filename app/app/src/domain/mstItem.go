package domain

import "time"

type MstItem struct {
  Id int
  ParentId int
  SubCategoryId int
  SizeGroupId int
  ColorGroupId int
  BrandId int
  ItemName string // char
  ItemDescription1 string // text
  ItemDescription2 string // text
  ItemDescription3 string // text
  ItemImg1 string // char
  ItemImg2 string // char
  ItemImg3 string // char
  DeleteFlag bool
  CreateDate time.Time
  UpdateDate time.Time
  DeleteDate time.Time
}
