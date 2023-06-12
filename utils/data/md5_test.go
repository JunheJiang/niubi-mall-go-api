package data

import "testing"

func TestMD5VCase1(t *testing.T) {
	str1 := []byte("原文1")
	str2 := []byte("原文2")
	result1 := MD5V(str1, 'a', 'b')
	result2 := MD5V(str2, 'a', 'b')
	if result2 == result1 {
		t.Error("TestMD5V失败、针对被篡改原文签名理应不同")
	}
}

func TestMD5VCase2(t *testing.T) {
	str1 := []byte("原文1")
	str2 := []byte("原文1")
	result1 := MD5V(str1, 'a', 'b')
	result2 := MD5V(str2, 'a', 'b')
	if result2 != result1 {
		t.Error("TestMD5V失败、针对相同原文签名应一致")
	}
}

func TestMD5VCase3(t *testing.T) {
	str1 := []byte("原文1")
	str2 := []byte("原文1")
	result1 := MD5V(str1, 'a', 'b')
	result2 := MD5V(str2, 'a')
	if result2 != result1 {
		t.Log("TestMD5V失败、针对相同原文不同盐值签名应不一致")
	}
}

func Test(t *testing.T) {
	t.Run("TestMD5VCase1", TestMD5VCase1)
	t.Run("TestMD5VCase2", TestMD5VCase2)
	t.Run("TestMD5VCase3", TestMD5VCase3)
}
