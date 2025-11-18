import CryptoJS from 'crypto-js'

// 密钥和偏移量建议使用每个项目独有的随机字符串
// 注意：前端加密只能防止明文存储，无法真正保护敏感数据
const SECRET_KEY = CryptoJS.enc.Utf8.parse('3333e6e143439161')
const SECRET_IV  = CryptoJS.enc.Utf8.parse('e3bbe7e3ba84431a')

// 加密
export const encrypt = (data) => {
  try {
    let dataStr = typeof data === 'object' ? JSON.stringify(data) : String(data)
    let encrypted = CryptoJS.AES.encrypt(
      CryptoJS.enc.Utf8.parse(dataStr),
      SECRET_KEY,
      {
        iv: SECRET_IV,
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7,
      }
    )
    return encrypted.toString()
  } catch (error) {
    console.error('Encryption failed:', error)
    return null
  }
}

// 解密
export const decrypt = (ciphertext) => {
  try {
    if (!ciphertext || typeof ciphertext !== 'string') {
      return null
    }

    let decrypted = CryptoJS.AES.decrypt(
      ciphertext,
      SECRET_KEY,
      {
        iv: SECRET_IV,
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7,
      }
    )

    const result = decrypted.toString(CryptoJS.enc.Utf8)

    // 如果解密结果为空，说明解密失败
    if (!result) {
      console.warn('Decryption failed: empty result')
      return null
    }

    return result
  } catch (error) {
    console.error('Decryption failed:', error)
    return null
  }
}

export function setEncryptedStorage(key, value) {
  try {
    const encrypted = encrypt(value)
    if (encrypted) {
      localStorage.setItem(key, encrypted)
      return true
    }
    return false
  } catch (error) {
    console.error('Failed to set encrypted storage:', error)
    return false
  }
}

export function getDecryptedStorage(key) {
  try {
    const raw = localStorage.getItem(key)
    if (!raw) return null

    const decrypted = decrypt(raw)

    // 尝试解析 JSON（如果原始数据是对象）
    if (decrypted) {
      try {
        return JSON.parse(decrypted)
      } catch {
        // 如果不是 JSON，直接返回字符串
        return decrypted
      }
    }

    return null
  } catch (error) {
    console.error('Failed to get decrypted storage:', error)
    return null
  }
}
