package respcode

// FDSAP
// CA-GABAY STATUS CODE
const (
	SUC_CODE_200     = "200"
	SUC_CODE_200_MSG = "Successful."

	SUC_CODE_201     = "201"
	SUC_CODE_201_MSG = "Successfully logged in."

	SUC_CODE_202     = "202"
	SUC_CODE_202_MSG = "Successfully logged out."

	SUC_CODE_203     = "203"
	SUC_CODE_203_MSG = "Successfully Inserted."

	SUC_CODE_204     = "204"
	SUC_CODE_204_MSG = "Successfully Updated."

	SUC_CODE_205     = "205"
	SUC_CODE_205_MSG = "Successfully Downloaded."

	SUC_CODE_206     = "206"
	SUC_CODE_206_MSG = "API is running."
)

const (
	ERR_CODE_100     = "100"
	ERR_CODE_100_MSG = "Validation Failed."

	ERR_CODE_101     = "101"
	ERR_CODE_101_MSG = "First Login: Password Reset Required."

	ERR_CODE_102     = "102"
	ERR_CODE_102_MSG = "Password expired."

	ERR_CODE_103     = "103"
	ERR_CODE_103_MSG = "Invalid password."

	ERR_CODE_104     = "104"
	ERR_CODE_104_MSG = "Invalid token"

	ERR_CODE_105     = "105"
	ERR_CODE_105_MSG = "User already logged in."

	ERR_CODE_106     = "106"
	ERR_CODE_106_MSG = "Password cannot be reused."

	ERR_CODE_107     = "107"
	ERR_CODE_107_MSG = "Invalid date"

	ERR_CODE_108     = "108"
	ERR_CODE_108_MSG = "Invalid phone number."

	ERR_CODE_109     = "109"
	ERR_CODE_109_MSG = "Invalid email address."

	ERR_CODE_110     = "110"
	ERR_CODE_110_MSG = "Expired token."

	ERR_CODE_111     = "111"
	ERR_CODE_111_MSG = "Token missing."

	ERR_CODE_111_MT     = "MT111"
	ERR_CODE_111_MT_MSG = "Malformed token."

	ERR_CODE_111_IT     = "IT111"
	ERR_CODE_111_IT_MSG = "Invalid token."

	ERR_CODE_112     = "112"
	ERR_CODE_112_MSG = "Invalid staff id."

	ERR_CODE_113     = "113"
	ERR_CODE_113_MSG = "Prohibited reused of password."

	ERR_CODE_114     = "114"
	ERR_CODE_114_MSG = "Required password reset."

	ERR_CODE_115     = "115"
	ERR_CODE_115_MSG = "Session id is missing."

	ERR_CODE_116     = "116"
	ERR_CODE_116_MSG = "Prohibited reused of email address"

	ERR_CODE_117     = "117"
	ERR_CODE_117_MSG = "Invalid OTP."

	ERR_CODE_118     = "118"
	ERR_CODE_118_MSG = "Expired OTP."

	ERR_CODE_300     = "300"
	ERR_CODE_300_MSG = "Internal servcer error."

	ERR_CODE_301     = "301"
	ERR_CODE_301_MSG = "Failed to parse data."

	ERR_CODE_302     = "302"
	ERR_CODE_302_MSG = "Failed to fetch data."

	ERR_CODE_303     = "303"
	ERR_CODE_303_MSG = "Failed to insert data."

	ERR_CODE_304     = "304"
	ERR_CODE_304_MSG = "Failed to update data."
)
