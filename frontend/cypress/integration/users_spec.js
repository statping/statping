/// <reference types="cypress" />

import "../support/commands"

context('Users Tests', () => {


  beforeEach(() => {
    cy.restoreLocalStorageCache();
  });

  afterEach(() => {
    cy.saveLocalStorageCache();
  });

  it('should Login', () => {
    cy.visit('/login')
    cy.get('#username').clear().type('admin')
    cy.get('#password').clear().type('admin')
    cy.get('button[type="submit"]').click()

    cy.get('.navbar-brand').should('contain', 'Statping')
    cy.getCookies()

    cy.getCookies().should('have.length', 1)
  })

  it('should goto users', () => {
    cy.visit('/dashboard/users')
    cy.get('#users_table > tr').should('have.length', 1)
    cy.get('#users_table > tr').eq(0).contains('admin')
  })

  it('should create new User', () => {
    cy.visit('/dashboard/users')
    cy.get('#username').clear().type('admin2')
    cy.get('#email').clear().type('info@admin.com')
    cy.get('#password').clear().type('password123')
    cy.get('#password_confirm').clear().type('password123')

    cy.get('button[type="submit"]').click()
  })

  // it('should confirm new user', () => {
  //   cy.visit('/dashboard/users')
  //   cy.get('#users_table > tr').should('have.length', 2)
  //   cy.get('#users_table > tr').eq(0).contains('admin')
  //   cy.get('#users_table > tr').eq(1).contains('admin2')
  // })

})
