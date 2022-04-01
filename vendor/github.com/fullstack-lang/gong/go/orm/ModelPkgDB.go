// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gong/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_ModelPkg_sql sql.NullBool
var dummy_ModelPkg_time time.Duration
var dummy_ModelPkg_sort sort.Float64Slice

// ModelPkgAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model modelpkgAPI
type ModelPkgAPI struct {
	gorm.Model

	models.ModelPkg

	// encoding of pointers
	ModelPkgPointersEnconding
}

// ModelPkgPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ModelPkgPointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// ModelPkgDB describes a modelpkg in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model modelpkgDB
type ModelPkgDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field modelpkgDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString

	// Declation for basic field modelpkgDB.PkgPath {{BasicKind}} (to be completed)
	PkgPath_Data sql.NullString
	// encoding of pointers
	ModelPkgPointersEnconding
}

// ModelPkgDBs arrays modelpkgDBs
// swagger:response modelpkgDBsResponse
type ModelPkgDBs []ModelPkgDB

// ModelPkgDBResponse provides response
// swagger:response modelpkgDBResponse
type ModelPkgDBResponse struct {
	ModelPkgDB
}

// ModelPkgWOP is a ModelPkg without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ModelPkgWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	PkgPath string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var ModelPkg_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"PkgPath",
}

type BackRepoModelPkgStruct struct {
	// stores ModelPkgDB according to their gorm ID
	Map_ModelPkgDBID_ModelPkgDB *map[uint]*ModelPkgDB

	// stores ModelPkgDB ID according to ModelPkg address
	Map_ModelPkgPtr_ModelPkgDBID *map[*models.ModelPkg]uint

	// stores ModelPkg according to their gorm ID
	Map_ModelPkgDBID_ModelPkgPtr *map[uint]*models.ModelPkg

	db *gorm.DB
}

func (backRepoModelPkg *BackRepoModelPkgStruct) GetDB() *gorm.DB {
	return backRepoModelPkg.db
}

// GetModelPkgDBFromModelPkgPtr is a handy function to access the back repo instance from the stage instance
func (backRepoModelPkg *BackRepoModelPkgStruct) GetModelPkgDBFromModelPkgPtr(modelpkg *models.ModelPkg) (modelpkgDB *ModelPkgDB) {
	id := (*backRepoModelPkg.Map_ModelPkgPtr_ModelPkgDBID)[modelpkg]
	modelpkgDB = (*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB)[id]
	return
}

// BackRepoModelPkg.Init set up the BackRepo of the ModelPkg
func (backRepoModelPkg *BackRepoModelPkgStruct) Init(db *gorm.DB) (Error error) {

	if backRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr != nil {
		err := errors.New("In Init, backRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr should be nil")
		return err
	}

	if backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB != nil {
		err := errors.New("In Init, backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB should be nil")
		return err
	}

	if backRepoModelPkg.Map_ModelPkgPtr_ModelPkgDBID != nil {
		err := errors.New("In Init, backRepoModelPkg.Map_ModelPkgPtr_ModelPkgDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.ModelPkg, 0)
	backRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr = &tmp

	tmpDB := make(map[uint]*ModelPkgDB, 0)
	backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB = &tmpDB

	tmpID := make(map[*models.ModelPkg]uint, 0)
	backRepoModelPkg.Map_ModelPkgPtr_ModelPkgDBID = &tmpID

	backRepoModelPkg.db = db
	return
}

// BackRepoModelPkg.CommitPhaseOne commits all staged instances of ModelPkg to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoModelPkg *BackRepoModelPkgStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for modelpkg := range stage.ModelPkgs {
		backRepoModelPkg.CommitPhaseOneInstance(modelpkg)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, modelpkg := range *backRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr {
		if _, ok := stage.ModelPkgs[modelpkg]; !ok {
			backRepoModelPkg.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoModelPkg.CommitDeleteInstance commits deletion of ModelPkg to the BackRepo
func (backRepoModelPkg *BackRepoModelPkgStruct) CommitDeleteInstance(id uint) (Error error) {

	modelpkg := (*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr)[id]

	// modelpkg is not staged anymore, remove modelpkgDB
	modelpkgDB := (*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB)[id]
	query := backRepoModelPkg.db.Unscoped().Delete(&modelpkgDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoModelPkg.Map_ModelPkgPtr_ModelPkgDBID), modelpkg)
	delete((*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr), id)
	delete((*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB), id)

	return
}

// BackRepoModelPkg.CommitPhaseOneInstance commits modelpkg staged instances of ModelPkg to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoModelPkg *BackRepoModelPkgStruct) CommitPhaseOneInstance(modelpkg *models.ModelPkg) (Error error) {

	// check if the modelpkg is not commited yet
	if _, ok := (*backRepoModelPkg.Map_ModelPkgPtr_ModelPkgDBID)[modelpkg]; ok {
		return
	}

	// initiate modelpkg
	var modelpkgDB ModelPkgDB
	modelpkgDB.CopyBasicFieldsFromModelPkg(modelpkg)

	query := backRepoModelPkg.db.Create(&modelpkgDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoModelPkg.Map_ModelPkgPtr_ModelPkgDBID)[modelpkg] = modelpkgDB.ID
	(*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr)[modelpkgDB.ID] = modelpkg
	(*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB)[modelpkgDB.ID] = &modelpkgDB

	return
}

// BackRepoModelPkg.CommitPhaseTwo commits all staged instances of ModelPkg to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoModelPkg *BackRepoModelPkgStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, modelpkg := range *backRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr {
		backRepoModelPkg.CommitPhaseTwoInstance(backRepo, idx, modelpkg)
	}

	return
}

// BackRepoModelPkg.CommitPhaseTwoInstance commits {{structname }} of models.ModelPkg to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoModelPkg *BackRepoModelPkgStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, modelpkg *models.ModelPkg) (Error error) {

	// fetch matching modelpkgDB
	if modelpkgDB, ok := (*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB)[idx]; ok {

		modelpkgDB.CopyBasicFieldsFromModelPkg(modelpkg)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoModelPkg.db.Save(&modelpkgDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown ModelPkg intance %s", modelpkg.Name))
		return err
	}

	return
}

// BackRepoModelPkg.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoModelPkg *BackRepoModelPkgStruct) CheckoutPhaseOne() (Error error) {

	modelpkgDBArray := make([]ModelPkgDB, 0)
	query := backRepoModelPkg.db.Find(&modelpkgDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	modelpkgInstancesToBeRemovedFromTheStage := make(map[*models.ModelPkg]struct{})
	for key, value := range models.Stage.ModelPkgs {
		modelpkgInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, modelpkgDB := range modelpkgDBArray {
		backRepoModelPkg.CheckoutPhaseOneInstance(&modelpkgDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		modelpkg, ok := (*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr)[modelpkgDB.ID]
		if ok {
			delete(modelpkgInstancesToBeRemovedFromTheStage, modelpkg)
		}
	}

	// remove from stage and back repo's 3 maps all modelpkgs that are not in the checkout
	for modelpkg := range modelpkgInstancesToBeRemovedFromTheStage {
		modelpkg.Unstage()

		// remove instance from the back repo 3 maps
		modelpkgID := (*backRepoModelPkg.Map_ModelPkgPtr_ModelPkgDBID)[modelpkg]
		delete((*backRepoModelPkg.Map_ModelPkgPtr_ModelPkgDBID), modelpkg)
		delete((*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB), modelpkgID)
		delete((*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr), modelpkgID)
	}

	return
}

// CheckoutPhaseOneInstance takes a modelpkgDB that has been found in the DB, updates the backRepo and stages the
// models version of the modelpkgDB
func (backRepoModelPkg *BackRepoModelPkgStruct) CheckoutPhaseOneInstance(modelpkgDB *ModelPkgDB) (Error error) {

	modelpkg, ok := (*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr)[modelpkgDB.ID]
	if !ok {
		modelpkg = new(models.ModelPkg)

		(*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr)[modelpkgDB.ID] = modelpkg
		(*backRepoModelPkg.Map_ModelPkgPtr_ModelPkgDBID)[modelpkg] = modelpkgDB.ID

		// append model store with the new element
		modelpkg.Name = modelpkgDB.Name_Data.String
		modelpkg.Stage()
	}
	modelpkgDB.CopyBasicFieldsToModelPkg(modelpkg)

	// preserve pointer to modelpkgDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ModelPkgDBID_ModelPkgDB)[modelpkgDB hold variable pointers
	modelpkgDB_Data := *modelpkgDB
	preservedPtrToModelPkg := &modelpkgDB_Data
	(*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB)[modelpkgDB.ID] = preservedPtrToModelPkg

	return
}

// BackRepoModelPkg.CheckoutPhaseTwo Checkouts all staged instances of ModelPkg to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoModelPkg *BackRepoModelPkgStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, modelpkgDB := range *backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB {
		backRepoModelPkg.CheckoutPhaseTwoInstance(backRepo, modelpkgDB)
	}
	return
}

// BackRepoModelPkg.CheckoutPhaseTwoInstance Checkouts staged instances of ModelPkg to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoModelPkg *BackRepoModelPkgStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, modelpkgDB *ModelPkgDB) (Error error) {

	modelpkg := (*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgPtr)[modelpkgDB.ID]
	_ = modelpkg // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitModelPkg allows commit of a single modelpkg (if already staged)
func (backRepo *BackRepoStruct) CommitModelPkg(modelpkg *models.ModelPkg) {
	backRepo.BackRepoModelPkg.CommitPhaseOneInstance(modelpkg)
	if id, ok := (*backRepo.BackRepoModelPkg.Map_ModelPkgPtr_ModelPkgDBID)[modelpkg]; ok {
		backRepo.BackRepoModelPkg.CommitPhaseTwoInstance(backRepo, id, modelpkg)
	}
}

// CommitModelPkg allows checkout of a single modelpkg (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutModelPkg(modelpkg *models.ModelPkg) {
	// check if the modelpkg is staged
	if _, ok := (*backRepo.BackRepoModelPkg.Map_ModelPkgPtr_ModelPkgDBID)[modelpkg]; ok {

		if id, ok := (*backRepo.BackRepoModelPkg.Map_ModelPkgPtr_ModelPkgDBID)[modelpkg]; ok {
			var modelpkgDB ModelPkgDB
			modelpkgDB.ID = id

			if err := backRepo.BackRepoModelPkg.db.First(&modelpkgDB, id).Error; err != nil {
				log.Panicln("CheckoutModelPkg : Problem with getting object with id:", id)
			}
			backRepo.BackRepoModelPkg.CheckoutPhaseOneInstance(&modelpkgDB)
			backRepo.BackRepoModelPkg.CheckoutPhaseTwoInstance(backRepo, &modelpkgDB)
		}
	}
}

// CopyBasicFieldsFromModelPkg
func (modelpkgDB *ModelPkgDB) CopyBasicFieldsFromModelPkg(modelpkg *models.ModelPkg) {
	// insertion point for fields commit

	modelpkgDB.Name_Data.String = modelpkg.Name
	modelpkgDB.Name_Data.Valid = true

	modelpkgDB.PkgPath_Data.String = modelpkg.PkgPath
	modelpkgDB.PkgPath_Data.Valid = true
}

// CopyBasicFieldsFromModelPkgWOP
func (modelpkgDB *ModelPkgDB) CopyBasicFieldsFromModelPkgWOP(modelpkg *ModelPkgWOP) {
	// insertion point for fields commit

	modelpkgDB.Name_Data.String = modelpkg.Name
	modelpkgDB.Name_Data.Valid = true

	modelpkgDB.PkgPath_Data.String = modelpkg.PkgPath
	modelpkgDB.PkgPath_Data.Valid = true
}

// CopyBasicFieldsToModelPkg
func (modelpkgDB *ModelPkgDB) CopyBasicFieldsToModelPkg(modelpkg *models.ModelPkg) {
	// insertion point for checkout of basic fields (back repo to stage)
	modelpkg.Name = modelpkgDB.Name_Data.String
	modelpkg.PkgPath = modelpkgDB.PkgPath_Data.String
}

// CopyBasicFieldsToModelPkgWOP
func (modelpkgDB *ModelPkgDB) CopyBasicFieldsToModelPkgWOP(modelpkg *ModelPkgWOP) {
	modelpkg.ID = int(modelpkgDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	modelpkg.Name = modelpkgDB.Name_Data.String
	modelpkg.PkgPath = modelpkgDB.PkgPath_Data.String
}

// Backup generates a json file from a slice of all ModelPkgDB instances in the backrepo
func (backRepoModelPkg *BackRepoModelPkgStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ModelPkgDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ModelPkgDB, 0)
	for _, modelpkgDB := range *backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB {
		forBackup = append(forBackup, modelpkgDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json ModelPkg ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json ModelPkg file", err.Error())
	}
}

// Backup generates a json file from a slice of all ModelPkgDB instances in the backrepo
func (backRepoModelPkg *BackRepoModelPkgStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ModelPkgDB, 0)
	for _, modelpkgDB := range *backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB {
		forBackup = append(forBackup, modelpkgDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("ModelPkg")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&ModelPkg_Fields, -1)
	for _, modelpkgDB := range forBackup {

		var modelpkgWOP ModelPkgWOP
		modelpkgDB.CopyBasicFieldsToModelPkgWOP(&modelpkgWOP)

		row := sh.AddRow()
		row.WriteStruct(&modelpkgWOP, -1)
	}
}

// RestoreXL from the "ModelPkg" sheet all ModelPkgDB instances
func (backRepoModelPkg *BackRepoModelPkgStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoModelPkgid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["ModelPkg"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoModelPkg.rowVisitorModelPkg)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoModelPkg *BackRepoModelPkgStruct) rowVisitorModelPkg(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var modelpkgWOP ModelPkgWOP
		row.ReadStruct(&modelpkgWOP)

		// add the unmarshalled struct to the stage
		modelpkgDB := new(ModelPkgDB)
		modelpkgDB.CopyBasicFieldsFromModelPkgWOP(&modelpkgWOP)

		modelpkgDB_ID_atBackupTime := modelpkgDB.ID
		modelpkgDB.ID = 0
		query := backRepoModelPkg.db.Create(modelpkgDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB)[modelpkgDB.ID] = modelpkgDB
		BackRepoModelPkgid_atBckpTime_newID[modelpkgDB_ID_atBackupTime] = modelpkgDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ModelPkgDB.json" in dirPath that stores an array
// of ModelPkgDB and stores it in the database
// the map BackRepoModelPkgid_atBckpTime_newID is updated accordingly
func (backRepoModelPkg *BackRepoModelPkgStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoModelPkgid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ModelPkgDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json ModelPkg file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ModelPkgDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ModelPkgDBID_ModelPkgDB
	for _, modelpkgDB := range forRestore {

		modelpkgDB_ID_atBackupTime := modelpkgDB.ID
		modelpkgDB.ID = 0
		query := backRepoModelPkg.db.Create(modelpkgDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB)[modelpkgDB.ID] = modelpkgDB
		BackRepoModelPkgid_atBckpTime_newID[modelpkgDB_ID_atBackupTime] = modelpkgDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json ModelPkg file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<ModelPkg>id_atBckpTime_newID
// to compute new index
func (backRepoModelPkg *BackRepoModelPkgStruct) RestorePhaseTwo() {

	for _, modelpkgDB := range *backRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB {

		// next line of code is to avert unused variable compilation error
		_ = modelpkgDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoModelPkg.db.Model(modelpkgDB).Updates(*modelpkgDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoModelPkgid_atBckpTime_newID map[uint]uint
