package models

const BackRepoTemplateCode = `// generated by genORMTranslation.go
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gorm.io/gorm"

	"{{PkgPathRoot}}/models"

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations{{` + string(rune(BackRepoPerStructDeclarations)) + `}}
	CommitFromBackNb uint // this ng is updated at the BackRepo level but also at the BackRepo<GongStruct> level

	PushFromFrontNb uint // records increments from push from front
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if models.Stage.OnInitCommitCallback != nil {
		models.Stage.OnInitCommitCallback.BeforeCommit(&models.Stage)
	}
	if models.Stage.OnInitCommitFromBackCallback != nil {
		models.Stage.OnInitCommitFromBackCallback.BeforeCommit(&models.Stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if models.Stage.OnInitCommitCallback != nil {
		models.Stage.OnInitCommitCallback.BeforeCommit(&models.Stage)
	}
	if models.Stage.OnInitCommitFromFrontCallback != nil {
		models.Stage.OnInitCommitFromFrontCallback.BeforeCommit(&models.Stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Init the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) init(db *gorm.DB) {
	// insertion point for per struct back repo declarations{{` + string(rune(BackRepoPerStructInits)) + `}}

	models.Stage.BackRepo = backRepo
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit{{` + string(rune(BackRepoPerStructPhaseOneCommits)) + `}}

	// insertion point for per struct back repo phase two commit{{` + string(rune(BackRepoPerStructPhaseTwoCommits)) + `}}

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit{{` + string(rune(BackRepoPerStructPhaseOneCheckouts)) + `}}

	// insertion point for per struct back repo phase two commit{{` + string(rune(BackRepoPerStructPhaseTwoCheckouts)) + `}}
}

var BackRepo BackRepoStruct

func GetLastCommitFromBackNb() uint {
	return BackRepo.GetLastCommitFromBackNb()
}

func GetLastPushFromFrontNb() uint {
	return BackRepo.GetLastPushFromFrontNb()
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup{{` + string(rune(BackRepoBackup)) + `}}
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup{{` + string(rune(BackRepoBackupXL)) + `}}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	models.Stage.Commit()
	models.Stage.Reset()
	models.Stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup{{` + string(rune(BackRepoRestorePhaseOne)) + `}}

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup{{` + string(rune(BackRepoRestorePhaseTwo)) + `}}

	models.Stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	models.Stage.Reset()

	// commit the cleaned stage
	models.Stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup{{` + string(rune(BackRepoRestoreXLPhaseOne)) + `}}

	// commit the restored stage
	models.Stage.Commit()
}
`

type BackRepoSubTemplateInsertion int

const (
	BackRepoPerStructDeclarations BackRepoSubTemplateInsertion = iota
	BackRepoPerStructInits
	BackRepoPerStructPhaseOneCommits
	BackRepoPerStructPhaseTwoCommits
	BackRepoPerStructPhaseOneCheckouts
	BackRepoPerStructPhaseTwoCheckouts
	BackRepoInitAndCommit
	BackRepoInitAndCheckout
	BackRepoCommit
	BackRepoCheckout
	BackRepoBackup
	BackRepoBackupXL
	BackRepoRestorePhaseOne
	BackRepoRestoreXLPhaseOne
	BackRepoRestorePhaseTwo
)

var BackRepoSubTemplate map[string]string = // new line
map[string]string{

	string(rune(BackRepoPerStructDeclarations)): `
	BackRepo{{Structname}} BackRepo{{Structname}}Struct
`,

	string(rune(BackRepoPerStructInits)): `
	backRepo.BackRepo{{Structname}}.Init(db)`,

	string(rune(BackRepoPerStructPhaseOneCommits)): `
	backRepo.BackRepo{{Structname}}.CommitPhaseOne(stage)`,

	string(rune(BackRepoPerStructPhaseTwoCommits)): `
	backRepo.BackRepo{{Structname}}.CommitPhaseTwo(backRepo)`,

	string(rune(BackRepoPerStructPhaseOneCheckouts)): `
	backRepo.BackRepo{{Structname}}.CheckoutPhaseOne()`,

	string(rune(BackRepoPerStructPhaseTwoCheckouts)): `
	backRepo.BackRepo{{Structname}}.CheckoutPhaseTwo(backRepo)`,

	string(rune(BackRepoInitAndCommit)): `
	map_{{Structname}}DBID_{{Structname}}DB = nil
	map_{{Structname}}Ptr_{{Structname}}DBID = nil
	map_{{Structname}}DBID_{{Structname}}Ptr = nil
	if err := BackRepo{{Structname}}Init(
		CreateMode,
		db); err != nil {
		return err
	}
`,

	string(rune(BackRepoInitAndCheckout)): `
	map_{{Structname}}DBID_{{Structname}}DB = nil
	map_{{Structname}}Ptr_{{Structname}}DBID = nil
	map_{{Structname}}DBID_{{Structname}}Ptr = nil
	if err := BackRepo{{Structname}}Init(
		CreateMode,
		db); err != nil {
		err := errors.New("AllORMToModels, CreateMode Translation of {{Structname}} failed")
		return err
	}
`,

	string(rune(BackRepoCheckout)): `
	if err := BackRepo{{Structname}}Init(
		UpdateMode,
		db); err != nil {
		err := errors.New("AllORMToModels, UpdateMode Translation of {{Structname}} failed")
		return err
	}
`,

	string(rune(BackRepoCommit)): `
	if err := BackRepo{{Structname}}Init(
		UpdateMode,
		db); err != nil {
		return err
	}
`,

	string(rune(BackRepoBackup)): `
	backRepo.BackRepo{{Structname}}.Backup(dirPath)`,

	string(rune(BackRepoBackupXL)): `
	backRepo.BackRepo{{Structname}}.BackupXL(file)`,

	string(rune(BackRepoRestorePhaseOne)): `
	backRepo.BackRepo{{Structname}}.RestorePhaseOne(dirPath)`,

	string(rune(BackRepoRestoreXLPhaseOne)): `
	backRepo.BackRepo{{Structname}}.RestoreXLPhaseOne(file)`,

	string(rune(BackRepoRestorePhaseTwo)): `
	backRepo.BackRepo{{Structname}}.RestorePhaseTwo()`,
}
