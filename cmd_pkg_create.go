package main

import (
	"fmt"
	// "io"
	// "io/ioutil"
	"os"
	"path/filepath"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func acvmfs_make_cmd_pkg_create() *commander.Command {
	cmd := &commander.Command{
		Run:       acvmfs_run_cmd_pkg_create,
		UsageLine: "pkg-create [options]",
		Short:     "create a new package tarball from CVMFS-based release",
		Long: `
pkg-create creates a new package tarball from CVMFS-based release.

ex:
 $ atl-cvmfs pkg-create -rel 17.2.10
 $ atl-cvmfs pkg-create -rel 17.2.10 -cmtcfg=i686-slc5-gcc43-opt
`,
		Flag: *flag.NewFlagSet("acvmfs-pkg-create", flag.ExitOnError),
		//CustomFlags: true,
	}
	cmd.Flag.Bool("q", false, "quiet. only print error and warning messages, all other output will be suppressed")
	cmd.Flag.String("rel", "", "release number to package up (e.g. 17.2.10)")
	cmd.Flag.String("cmtcfg", "", "CMTCONFIG to use (default=${CMTCONFIG})")
	cmd.Flag.String("cvmfsdir", "/cvmfs/atlas.cern.ch/repo/sw/software", "top directory under which all releases are located")
	cmd.Flag.String("outdir", ".", "directory where to put the package tarball")
	cmd.Flag.Bool("with-dbrelease", true, "include DBRelease in package tarball")
	return cmd
}

func acvmfs_run_cmd_pkg_create(cmd *commander.Command, args []string) {
	var err error
	n := "acvmfs-" + cmd.Name()

	switch len(args) {
	case 0:
		// ok
	default:
		err = fmt.Errorf("%s: does not take any argument\n", n)
		handle_err(err)
	}

	quiet := cmd.Flag.Lookup("q").Value.Get().(bool)
	release := cmd.Flag.Lookup("rel").Value.Get().(string)
	cmtcfg := cmd.Flag.Lookup("cmtcfg").Value.Get().(string)
	cvmfsdir := cmd.Flag.Lookup("cvmfsdir").Value.Get().(string)
	outdir := cmd.Flag.Lookup("outdir").Value.Get().(string)
	withdbrelease := cmd.Flag.Lookup("with-dbrelease").Value.Get().(bool)

	if release == "" {
		err = fmt.Errorf("you need to give a release number to package up (e.g: 17.2.10)")
		handle_err(err)
	}

	if cvmfsdir == "" {
		err = fmt.Errorf("you need to give a path to CVMFS top directory (e.g: /cvmfs/atlas.cern.ch/repo/sw/software)")
		handle_err(err)
	}

	if cmtcfg == "" {
		cmtcfg = os.Getenv("CMTCONFIG")
		if cmtcfg == "" {
			cmtcfg = "i686-slc5-gcc43-opt"
		}
	}

	if !quiet {
		fmt.Printf("%s: creating package [athena-%s-%s]...\n",
			n, release, cmtcfg)
	}

	if !path_exists(cvmfsdir) {
		err = fmt.Errorf("no such CVMFS directory [%s]", cvmfsdir)
		handle_err(err)
	}

	swdir := filepath.Join(cvmfsdir, cmtcfg, release)
	if !path_exists(swdir) {
		err = fmt.Errorf("no such atlas s/w directory [%s]", swdir)
		handle_err(err)
	}

	// handle env. vars in case the shell didn't do it already
	outdir = os.ExpandEnv(outdir)
	outdir, err = filepath.Abs(outdir)
	handle_err(err)

	// workdir, err := ioutil.TempDir("", "atl-cvmfs-")
	// handle_err(err)
	// defer os.RemoveAll(workdir)

	pkgname := fmt.Sprintf("athena-%s-%s.tar.gz", release, cmtcfg)
	targ := filepath.Join(outdir, pkgname)

	if !quiet {
		fmt.Printf("%s: swdir=   %s\n", n, swdir)
		fmt.Printf("%s: outdir=  %s\n", n, outdir)
		fmt.Printf("%s: pkgname= %s\n", n, pkgname)
	}

	if !path_exists(outdir) {
		err = os.MkdirAll(outdir, 0700)
		handle_err(err)
	}

	extra_args := []string{}
	if !withdbrelease {
		extra_args = append(extra_args, "--exclude=DBRelease")
	}
	err = _tar_gz(targ, swdir, extra_args)
	handle_err(err)

	// src, err := os.Open(targ)
	// handle_err(err)
	// defer src.Close()

	// dst, err := os.Create(filepath.Join(outdir, pkgname))
	// handle_err(err)
	// defer dst.Close()

	// _, err = io.Copy(dst, src)
	// handle_err(err)
	// err = dst.Sync()
	// handle_err(err)
	// err = dst.Close()
	// handle_err(err)

	if !quiet {
		fmt.Printf("%s: creating package [athena-%s-%s]... [OK]\n",
			n, release, cmtcfg)
	}
	return
}

// EOF
