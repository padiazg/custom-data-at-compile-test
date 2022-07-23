package main

var (
	buildTime  string
	commitHash string
	version    string
)

func GetVersion() string {
	if len(version) > 0 {
		return version
	}
	return "Dev"
}

func GetCommitHash() string {
	if len(commitHash) > 0 {
		return commitHash
	}
	return ""
}

func GetBuildTime() string {
	if len(buildTime) > 0 {
		return buildTime
	}

	return ""
}
