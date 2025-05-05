package io

import (
    "github.com/go-git/go-billy/v5"
    "github.com/msisdev/dotato/pkg/config"
    "github.com/msisdev/dotato/pkg/dotato"
    "github.com/msisdev/dotato/pkg/state"
)

func ImportFile(
    fs billy.Filesystem,
    dtt *dotato.Dotato,
    pre dotato.PreviewImportFile,
) error {
    if pre.Equal {
        return nil
    }

    // Paths
    var (
        dotPath string
        dttPath = pre.Dtt.Abs()
    )
    if pre.DotReal != nil {
        dotPath = pre.DotReal.Abs()
    } else {
        dotPath = pre.Dot.Abs()
    }

    // Make dotato directory
    err := fs.MkdirAll(pre.Dtt.Parent().Abs(), 0755)
    if err != nil {
        return err
    }

    // Copy file
    err = dtt.CopyFile(dotPath, dttPath)
    if err != nil {
        return err
    }

    // Write history
    err = dtt.PutHistory(state.History{
        DotPath: dotPath,
        DttPath: dttPath,
        Mode:    config.ModeFile,
    })
    if err != nil {
        return err
    }

    return nil
}