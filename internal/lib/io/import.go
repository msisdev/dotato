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
    pre dotato.Preview,
) error {
  if pre.DttOp == dotato.FileOpNone {
    return nil
  }

  var (
    dotabs = pre.Dot.Path.Abs()
    dttabs = pre.Dtt.Path.Abs()
  )

  if pre.DttOp == dotato.FileOpOverwrite {
    // Remove dtt
    err := fs.Remove(dttabs)
    if err != nil {
      return err
    }
  } else {
    // Make directory
    err := fs.MkdirAll(pre.Dtt.Path.Parent().Abs(), 0755)
    if err != nil {
      return err
    }
  }

  // Copy file
  err := dtt.CreateAndCopyFile(dotabs, dttabs)
  if err != nil {
    return err
  }

  // Write history
  err = dtt.PutHistory(state.History{
    DotPath: dotabs,
    DttPath: dttabs,
    Mode:    config.ModeFile,
  })
  if err != nil {
    return err
  }

  return nil
}

func ImportLink(
  fs billy.Filesystem,
  dtt *dotato.Dotato,
  pre dotato.Preview,
) error {
  if pre.DotOp == dotato.FileOpNone && 
    pre.DttOp == dotato.FileOpNone {
    return nil
  }

  var (
    dotabs = pre.Dot.Path.Abs()
    dttabs = pre.Dtt.Path.Abs()
  )

  // Create dtt first
  if pre.DttOp != dotato.FileOpNone {
    if pre.DttOp == dotato.FileOpOverwrite {
      // Remove dtt
      err := fs.Remove(dttabs)
      if err != nil {
        return err
      }
    } else {
      // Create directory
      err := fs.MkdirAll(pre.Dtt.Path.Parent().Abs(), 0755)
      if err != nil {
        return err
      }
    }

    // Create file
    err := dtt.CreateAndCopyFile(dotabs, dttabs)
    if err != nil {
      return err
    }
  }

  // Create dot after
  if pre.DotOp != dotato.FileOpNone {
    if pre.DotOp == dotato.FileOpOverwrite {
      // Remove dot
      err := fs.Remove(dotabs)
      if err != nil {
        return err
      }
    } else {
      // Create directory
      err := fs.MkdirAll(pre.Dot.Path.Parent().Abs(), 0755)
      if err != nil {
        return err
      }
    }

    // Create link
    err := fs.Symlink(dttabs, dotabs)
    if err != nil {
      return err
    }
  }

  // Write history
  err := dtt.PutHistory(state.History{
    DotPath:  dotabs,
    DttPath:  dttabs,
    Mode:     config.ModeLink,
  })
  if err != nil {
    return err
  }

  return nil
}

func ExportFile(
  fs billy.Filesystem,
  dtt *dotato.Dotato,
  pre dotato.Preview,
) error {
  if pre.DotOp == dotato.FileOpNone {
    return nil
  }

  var (
    dotabs = pre.Dot.Path.Abs()
    dttabs = pre.Dtt.Path.Abs()
  )

  if pre.DotOp == dotato.FileOpOverwrite {
    // Do nothing
  } else {
    // Make directory
    err := fs.MkdirAll(pre.Dot.Path.Parent().Abs(), 0755)
    if err != nil {
      return err
    }
  }

  // Copy file
  err := dtt.CreateAndCopyFile(dttabs, dotabs)
  if err != nil {
    return err
  }

  // Write history
  err = dtt.PutHistory(state.History{
    DotPath: dotabs,
    DttPath: dttabs,
    Mode:    config.ModeFile,
  })
  if err != nil {
    return err
  }

  return nil
}

func ExportLink(
  fs billy.Filesystem,
  dtt *dotato.Dotato,
  pre dotato.Preview,
) error {
  if pre.DotOp == dotato.FileOpNone {
    return nil
  }

  var (
    dotabs = pre.Dot.Path.Abs()
    dttabs = pre.Dtt.Path.Abs()
  )

  if pre.DotOp == dotato.FileOpOverwrite {
    // Remove dot
    err := fs.Remove(dotabs)
    if err != nil {
      return err
    }
  } else {
    // Make directory
    err := fs.MkdirAll(pre.Dot.Path.Parent().Abs(), 0755)
    if err != nil {
      return err
    }
  }

  // Create link
  err := fs.Symlink(dttabs, dotabs)
  if err != nil {
    return err
  }

  // Write history
  err = dtt.PutHistory(state.History{
    DotPath: dotabs,
    DttPath: dttabs,
    Mode:    config.ModeLink,
  })
  if err != nil {
    return err
  }

  return nil
}
