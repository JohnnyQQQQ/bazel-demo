load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_binary", "go_library")

gazelle(
    name = "gazelle",
    prefix = "my-service",
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "my-service",
    visibility = ["//visibility:private"],
    deps = ["@com_github_gin_gonic_gin//:go_default_library"],
)

go_binary(
    name = "my-service",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
