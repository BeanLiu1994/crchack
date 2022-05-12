cmake_minimum_required(VERSION 3.0.0)
project(crchack VERSION 0.1.0 LANGUAGES C)

add_executable(crchack bigint.c crc.c crchack.c forge.c)

IF (WIN32)
find_program(GO_COMPILER NAMES go.exe PATHS PATH)
ELSE()
find_program(GO_COMPILER NAMES go PATHS PATH)
ENDIF()

add_custom_target(app
        ALL
        WORKING_DIRECTORY "${PROJECT_SOURCE_DIR}/app/"
        COMMAND ${GO_COMPILER} "build" "-o" "${PROJECT_BINARY_DIR}/app" "main.go"
        DEPENDS crchack
        BYPRODUCTS "${PROJECT_BINARY_DIR}/app")

install(TARGETS crchack
        RUNTIME DESTINATION ${PROJECT_SOURCE_DIR}/bin)
install(FILES "${PROJECT_BINARY_DIR}/app" DESTINATION ${PROJECT_SOURCE_DIR}/bin)